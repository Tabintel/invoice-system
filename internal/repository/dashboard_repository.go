package repository

import (
    "context"
    "github.com/Tabintel/invoice-system/internal/ent"
    //"github.com/Tabintel/invoice-system/ent/invoice"
)

type DashboardRepository struct {
    client *ent.Client
}

func NewDashboardRepository(client *ent.Client) *DashboardRepository {
    return &DashboardRepository{client: client}
}

func (r *DashboardRepository) GetStatistics(ctx context.Context) (*DashboardStats, error) {
    stats := &DashboardStats{}
    
    // Total Paid
    paid, err := r.client.Invoice.Query().
        Where(invoice.StatusEQ("paid")).
        Aggregate(ent.Count(), ent.Sum(invoice.FieldAmount)).
        First(ctx)
    if err != nil {
        return nil, err
    }
    stats.TotalPaid = paid[0]
    stats.TotalPaidAmount = paid[1]

    // Total Overdue
    overdue, err := r.client.Invoice.Query().
        Where(invoice.StatusEQ("overdue")).
        Aggregate(ent.Count(), ent.Sum(invoice.FieldAmount)).
        First(ctx)
    if err != nil {
        return nil, err
    }
    stats.TotalOverdue = overdue[0]
    stats.TotalOverdueAmount = overdue[1]

    return stats, nil
}
