package service

import (
    "context"
    "fmt"
    "time"
    "github.com/Tabintel/invoice-system/internal/repository"
)

type InvoiceService struct {
    repo               *repository.Repository
    userRepo           *repository.Repository
    notificationService *NotificationService
}

func NewInvoiceService(repo *repository.Repository, userRepo *repository.Repository, notificationService *NotificationService) *InvoiceService {
    return &InvoiceService{repo: repo, userRepo: userRepo, notificationService: notificationService}
}

type CreateInvoiceRequest struct {
    CustomerDetails struct {
        Name    string `json:"name"`
        Email   string `json:"email"`
        Phone   string `json:"phone"`
    } `json:"customer_details"`
    Items []struct {
        Description string  `json:"description"`
        Quantity    int     `json:"quantity"`
        Rate       float64 `json:"rate"`
    } `json:"items"`
    Currency    string    `json:"currency"`
    IssueDate   time.Time `json:"issue_date"`
    DueDate     time.Time `json:"due_date"`
    Notes       string    `json:"notes"`
}

func (s *InvoiceService) CreateInvoice(ctx context.Context, req *CreateInvoiceRequest, senderID int64) (*repository.Invoice, error) {
    // Start transaction
    tx, err := s.repo.BeginTx(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed to start transaction: %w", err)
    }
    defer tx.Rollback()

    // Calculate total amount
    var totalAmount float64
    for _, item := range req.Items {
        totalAmount += float64(item.Quantity) * item.Rate
    }

    // Create or get customer
    customerID, err := s.userRepo.CreateOrGetCustomer(ctx, req.CustomerDetails)
    if err != nil {
        return nil, fmt.Errorf("failed to process customer: %w", err)
    }

    // Create invoice
    invoice := &repository.Invoice{
        ReferenceNumber: generateReferenceNumber(),
        SenderID:       senderID,
        CustomerID:     customerID,
        Amount:         totalAmount,
        Currency:       req.Currency,
        Status:         "draft",
        IssueDate:      req.IssueDate,
        DueDate:        req.DueDate,
        Notes:          req.Notes,
    }

    if err := s.repo.CreateTx(ctx, tx, invoice); err != nil {
        return nil, fmt.Errorf("failed to create invoice: %w", err)
    }

    // Create invoice items
    for _, item := range req.Items {
        if err := s.repo.CreateInvoiceItem(ctx, tx, invoice.ID, item); err != nil {
            return nil, fmt.Errorf("failed to create invoice item: %w", err)
        }
    }

    if err := tx.Commit(); err != nil {
        return nil, fmt.Errorf("failed to commit transaction: %w", err)
    }

    return invoice, nil
}
func generateReferenceNumber() string {
    return fmt.Sprintf("INV-%d", time.Now().Unix())
}

func (s *InvoiceService) UpdateInvoiceStatus(ctx context.Context, invoiceID int64, userID int64, status string) error {
    // Verify invoice ownership
    invoice, err := s.repo.GetByID(ctx, invoiceID)
    if err != nil {
        return fmt.Errorf("failed to get invoice: %w", err)
    }
    if invoice.SenderID != userID {
        return fmt.Errorf("unauthorized: invoice belongs to different user")
    }

    // Validate status transition
    if !isValidStatusTransition(invoice.Status, status) {
        return fmt.Errorf("invalid status transition from %s to %s", invoice.Status, status)
    }

    tx, err := s.repo.BeginTx(ctx)
    if err != nil {
        return fmt.Errorf("failed to start transaction: %w", err)
    }
    defer tx.Rollback()

    // Update status
    invoice.Status = status
    invoice.UpdatedAt = time.Now()
    
    if err := s.repo.UpdateInvoiceTx(ctx, tx, invoice); err != nil {
        return fmt.Errorf("failed to update invoice: %w", err)
    }

    // Log activity
    activity := &repository.Activity{
        UserID:    userID,
        InvoiceID: invoiceID,
        Action:    fmt.Sprintf("status_updated_to_%s", status),
        Details:   fmt.Sprintf("Invoice status updated to %s", status),
    }
    
    if err := s.repo.LogActivityTx(ctx, tx, activity); err != nil {
        return fmt.Errorf("failed to log activity: %w", err)
    }

    if err := tx.Commit(); err != nil {
        return fmt.Errorf("failed to commit transaction: %w", err)
    }

    // Send notification after successful update
    go s.notificationService.SendInvoiceStatusUpdate(invoice, status)

    return nil
}

func isValidStatusTransition(from, to string) bool {
    transitions := map[string][]string{
        "draft":   {"pending"},
        "pending": {"paid", "overdue"},
        "paid":    {},
        "overdue": {"paid"},
    }
    
    allowedTransitions, exists := transitions[from]
    if !exists {
        return false
    }
    
    for _, allowed := range allowedTransitions {
        if allowed == to {
            return true
        }
    }
    return false
}
