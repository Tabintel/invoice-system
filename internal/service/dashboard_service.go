package service

import (
    "context"
    "github.com/Tabintel/invoice-system/internal/repository"
)

type DashboardService struct {
    repo *repository.Repository
}

func NewDashboardService(repo *repository.Repository) *DashboardService {
    return &DashboardService{repo: repo}
}

type DashboardResponse struct {
    Stats struct {
        Paid    StatInfo `json:"paid"`
        Overdue StatInfo `json:"overdue"`
        Draft   StatInfo `json:"draft"`
        Unpaid  StatInfo `json:"unpaid"`
    } `json:"stats"`
    RecentInvoices []InvoiceInfo    `json:"recent_invoices"`
    Activities     []ActivityInfo    `json:"activities"`
}

type StatInfo struct {
    Count int     `json:"count"`
    Total float64 `json:"total"`
}

func (s *DashboardService) GetDashboard(ctx context.Context, userID int64) (*DashboardResponse, error) {
    stats, err := s.invoiceRepo.GetDashboardStats(ctx, userID)
    if err != nil {
        return nil, err
    }

    recentInvoices, err := s.invoiceRepo.GetRecentInvoices(ctx, userID, 5)
    if err != nil {
        return nil, err
    }

    activities, err := s.invoiceRepo.GetRecentActivities(ctx, userID, 5)
    if err != nil {
        return nil, err
    }

    return &DashboardResponse{
        Stats: mapStatsToResponse(stats),
        RecentInvoices: mapInvoicesToResponse(recentInvoices),
        Activities: mapActivitiesToResponse(activities),
    }, nil
}
