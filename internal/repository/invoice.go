package repository

import (
    "context"
    "database/sql"
    "time"
)

type Invoice struct {
    ID              int64     `json:"id"`
    ReferenceNumber string    `json:"reference_number"`
    SenderID        int64     `json:"sender_id"`
    CustomerID      int64     `json:"customer_id"`
    Amount          float64   `json:"amount"`
    Currency        string    `json:"currency"`
    Status          string    `json:"status"`
    IssueDate       time.Time `json:"issue_date"`
    DueDate         time.Time `json:"due_date"`
    Notes           string    `json:"notes"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}

type InvoiceRepository struct {
    db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *InvoiceRepository {
    return &InvoiceRepository{db: db}
}

func (r *InvoiceRepository) Create(ctx context.Context, inv *Invoice) error {
    query := `
        INSERT INTO invoices (reference_number, sender_id, customer_id, amount, 
            currency, status, issue_date, due_date, notes)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING id, created_at, updated_at`

    return r.db.QueryRowContext(ctx, query,
        inv.ReferenceNumber, inv.SenderID, inv.CustomerID, inv.Amount,
        inv.Currency, inv.Status, inv.IssueDate, inv.DueDate, inv.Notes,
    ).Scan(&inv.ID, &inv.CreatedAt, &inv.UpdatedAt)
}

func (r *InvoiceRepository) GetByID(ctx context.Context, id int64) (*Invoice, error) {
    inv := &Invoice{}
    query := `SELECT * FROM invoices WHERE id = $1`
    
    err := r.db.QueryRowContext(ctx, query, id).Scan(
        &inv.ID, &inv.ReferenceNumber, &inv.SenderID, &inv.CustomerID,
        &inv.Amount, &inv.Currency, &inv.Status, &inv.IssueDate,
        &inv.DueDate, &inv.Notes, &inv.CreatedAt, &inv.UpdatedAt,
    )
    if err != nil {
        return nil, err
    }
    return inv, nil
}

func (r *InvoiceRepository) GetDashboardStats(ctx context.Context, userID int64) (map[string]DashboardStat, error) {
    query := `
        SELECT 
            status,
            COUNT(*) as count,
            COALESCE(SUM(amount), 0) as total_amount
        FROM invoices
        WHERE sender_id = $1
        GROUP BY status`

    rows, err := r.db.QueryContext(ctx, query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    stats := make(map[string]DashboardStat)
    for rows.Next() {
        var stat DashboardStat
        var status string
        if err := rows.Scan(&status, &stat.Count, &stat.TotalAmount); err != nil {
            return nil, err
        }
        stats[status] = stat
    }
    return stats, nil
}

type DashboardStat struct {
    Count       int     `json:"count"`
    TotalAmount float64 `json:"total_amount"`
}
