package services

import (
    "context"
    "time"
    "fmt"
    "crypto/rand"
    "encoding/base64"
    "os"
    "github.com/Tabintel/invoice-system/ent"
    "github.com/Tabintel/invoice-system/ent/invoice"
    
)

type InvoiceService struct {
    db *ent.Client
}

func NewInvoiceService(db *ent.Client) *InvoiceService {
    return &InvoiceService{db: db}
}

type CreateInvoiceInput struct {
    CustomerID     int     `json:"customer_id"`
    DueDate       time.Time `json:"due_date"`
    Currency      string   `json:"currency"`
    Items         []InvoiceItemInput `json:"items"`
}

type InvoiceItemInput struct {
    Description string  `json:"description"`
    Quantity    int     `json:"quantity"`
    Rate        float64 `json:"rate"`
}

func calculateTotalAmount(items []InvoiceItemInput) float64 {
    var total float64
    for _, item := range items {
        total += float64(item.Quantity) * item.Rate
    }
    return total
}

func generateReferenceNumber() string {
    timestamp := time.Now()
    return fmt.Sprintf("INV-%s-%d", 
        timestamp.Format("20060102"), 
        timestamp.UnixNano() % 10000)
}

func (s *InvoiceService) CreateInvoice(ctx context.Context, input CreateInvoiceInput) (*ent.Invoice, error) {
    totalAmount := calculateTotalAmount(input.Items)
    
    return s.db.Invoice.Create().
        SetReferenceNumber(generateReferenceNumber()).
        SetStatus("draft").
        SetIssueDate(time.Now()).
        SetDueDate(input.DueDate).
        SetCurrency(input.Currency).
        SetTotalAmount(totalAmount).
        Save(ctx)
}

// Adding new types to list invoices
type ListInvoicesInput struct {
    Page     int    `json:"page"`
    PageSize int    `json:"page_size"`
    Status   string `json:"status"`
    SortBy   string `json:"sort_by"`
}

type InvoiceStats struct {
    TotalPaid    int `json:"total_paid"`
    TotalOverdue int `json:"total_overdue"`
    TotalDraft   int `json:"total_draft"`
    TotalUnpaid  int `json:"total_unpaid"`
}

type UpdateInvoiceStatusInput struct {
    Status string `json:"status"`
}


// Added these methods to InvoiceService
func (s *InvoiceService) ListInvoices(ctx context.Context, input ListInvoicesInput) ([]*ent.Invoice, error) {
    query := s.db.Invoice.Query()
    
    if input.Status != "" {
        query = query.Where(invoice.Status(input.Status))
    }
    
    // Default sorting by created_at desc
    query = query.Order(ent.Desc(invoice.FieldCreatedAt))
    
    if input.PageSize == 0 {
        input.PageSize = 10
    }
    
    offset := (input.Page - 1) * input.PageSize
    return query.Limit(input.PageSize).Offset(offset).All(ctx)
}

func (s *InvoiceService) GetInvoiceStats(ctx context.Context) (*InvoiceStats, error) {
    stats := &InvoiceStats{}
    
    var err error
    stats.TotalPaid, err = s.db.Invoice.Query().Where(invoice.Status("paid")).Count(ctx)
    if err != nil {
        return nil, err
    }
    
    stats.TotalOverdue, err = s.db.Invoice.Query().Where(invoice.Status("overdue")).Count(ctx)
    if err != nil {
        return nil, err
    }
    
    stats.TotalDraft, err = s.db.Invoice.Query().Where(invoice.Status("draft")).Count(ctx)
    if err != nil {
        return nil, err
    }
    
    stats.TotalUnpaid, err = s.db.Invoice.Query().Where(invoice.Status("unpaid")).Count(ctx)
    if err != nil {
        return nil, err
    }
    
    return stats, nil
}
// InvoiceService
func (s *InvoiceService) UpdateStatus(ctx context.Context, id int, input UpdateInvoiceStatusInput) (*ent.Invoice, error) {
    return s.db.Invoice.UpdateOneID(id).
        SetStatus(input.Status).
        Save(ctx)
}

func (s *InvoiceService) GetInvoice(ctx context.Context, id int) (*ent.Invoice, error) {
    return s.db.Invoice.Get(ctx, id)
}

type ShareableLinkResponse struct {
    URL       string    `json:"url"`
    ExpiresAt time.Time `json:"expires_at"`
}

func (s *InvoiceService) GenerateShareableLink(ctx context.Context, id int) (*ShareableLinkResponse, error) {
    invoice, err := s.GetInvoice(ctx, id)
    if err != nil {
        return nil, err
    }
    
    token := generateSecureToken()
    expiryTime := time.Now().Add(7 * 24 * time.Hour)
    
    // Store token with invoice
    _, err = s.db.Invoice.UpdateOne(invoice).
        SetShareToken(token).
        SetShareExpiry(expiryTime).
        Save(ctx)
    if err != nil {
        return nil, err
    }
    
    shareableURL := fmt.Sprintf("%s/public/invoices/%s", os.Getenv("BASE_URL"), token)
    
    return &ShareableLinkResponse{
        URL:       shareableURL,
        ExpiresAt: expiryTime,
    }, nil
}
func generateSecureToken() string {
    b := make([]byte, 32)
    rand.Read(b)
    return base64.URLEncoding.EncodeToString(b)
}

func (s *InvoiceService) GetInvoiceByShareToken(ctx context.Context, token string) (*ent.Invoice, error) {
    return s.db.Invoice.Query().
        Where(invoice.ShareToken(token)).
        Only(ctx)
}
