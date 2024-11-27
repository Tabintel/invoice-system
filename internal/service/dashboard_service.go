package service

import (
    "context"
    "github.com/Tabintel/invoice-system/internal/repository"
)

type DashboardService struct {
    repo *repository.Repository
}

type DashboardResponse struct {
    Stats          map[string]interface{} `json:"stats"`
    RecentInvoices []repository.InvoiceInfo  `json:"recent_invoices"`
    RecentActivity []repository.ActivityInfo `json:"recent_activity"`
}

func NewDashboardService(repo *repository.Repository) *DashboardService {
    return &DashboardService{repo: repo}
}

func (s *DashboardService) GetDashboard(ctx context.Context, userID int64) (*DashboardResponse, error) {
    stats, err := s.repo.GetDashboardStats(ctx, userID)
    if err != nil {
        return nil, err
    }

    invoices, err := s.repo.GetRecentInvoices(ctx, userID, 5)
    if err != nil {
        return nil, err
    }

    activities, err := s.repo.GetRecentActivities(ctx, userID, 5)
    if err != nil {
        return nil, err
    }

    return &DashboardResponse{
        Stats:          stats,
        RecentInvoices: mapInvoicesToResponse(invoices),
        RecentActivity: mapActivitiesToResponse(activities),
    }, nil
}

func mapInvoicesToResponse(invoices []repository.Invoice) []repository.InvoiceInfo {
    var response []repository.InvoiceInfo
    for _, inv := range invoices {
        response = append(response, repository.InvoiceInfo{
            ID:              inv.ID,
            ReferenceNumber: inv.ReferenceNumber,
            Amount:          inv.Amount,
            Status:          inv.Status,
            DueDate:         inv.DueDate,
        })
    }
    return response
}

func mapActivitiesToResponse(activities []repository.Activity) []repository.ActivityInfo {
    var response []repository.ActivityInfo
    for _, act := range activities {
        response = append(response, repository.ActivityInfo{
            ID:            act.ID,
            Action:        act.Action,
            Details:       act.Details,
            InvoiceNumber: act.InvoiceNumber,
            Timestamp:     act.CreatedAt,
        })
    }
    return response
}