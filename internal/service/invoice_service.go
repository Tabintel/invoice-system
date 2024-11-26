package service

import (
    "context"
    "fmt"
    "time"
    "github.com/Tabintel/invoice-system/ent"
    "github.com/jung-kurt/gofpdf"
)

type InvoiceService struct {
    repo            *repository.InvoiceRepository
    activityRepo    *repository.ActivityRepository
    dashboardRepo   *repository.DashboardRepository
}

func (s *InvoiceService) GetDashboardStats(ctx context.Context) (*DashboardStats, error) {
    return s.dashboardRepo.GetStatistics(ctx)
}

func (s *InvoiceService) GeneratePDF(ctx context.Context, invoiceID int) ([]byte, error) {
    invoice, err := s.repo.GetByID(ctx, invoiceID)
    if err != nil {
        return nil, err
    }

    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    // Add invoice details to PDF
    pdf.SetFont("Arial", "B", 16)
    pdf.Cell(40, 10, fmt.Sprintf("Invoice #%s", invoice.ReferenceNumber))
    
    return pdf.OutputBytes()
}

func (s *InvoiceService) GenerateShareableLink(ctx context.Context, invoiceID int) (string, error) {
    // Generate secure token
    token := generateSecureToken()
    
    // Store token with expiry
    err := s.repo.StoreShareableLink(ctx, invoiceID, token)
    if err != nil {
        return "", err
    }
    
    return fmt.Sprintf("/invoice/share/%s", token), nil
}

func (s *InvoiceService) UpdateStatus(ctx context.Context, invoiceID int, newStatus string) error {
    err := s.repo.UpdateStatus(ctx, invoiceID, newStatus)
    if err != nil {
        return err
    }
    
    // Log activity
    s.activityRepo.LogActivity(ctx, &ent.ActivityLog{
        ActionType: "status_update",
        Description: fmt.Sprintf("Invoice status updated to %s", newStatus),
    })
    
    return nil
}
