package repository

import (
    "context"
    "database/sql"
    "time"
)

type InvoiceRepository struct {
    db *sql.DB
}

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
    Items           []InvoiceItem `json:"items,omitempty"`
}

type InvoiceItem struct {
    ID          int64   `json:"id"`
    InvoiceID   int64   `json:"invoice_id"`
    Description string  `json:"description"`
    Quantity    int     `json:"quantity"`
    Rate        float64 `json:"rate"`
    Amount      float64 `json:"amount"`
}

func NewInvoiceRepository(db *sql.DB) *InvoiceRepository {
    return &InvoiceRepository{db: db}
}

func (r *InvoiceRepository) BeginTx(ctx context.Context) (*sql.Tx, error) {
    return r.db.BeginTx(ctx, nil)
}

func (r *InvoiceRepository) CreateTx(ctx context.Context, tx *sql.Tx, inv *Invoice) error {
    query := `
        INSERT INTO invoices (reference_number, sender_id, customer_id, amount, 
            currency, status, issue_date, due_date, notes)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING id, created_at, updated_at`

    return tx.QueryRowContext(ctx, query,
        inv.ReferenceNumber, inv.SenderID, inv.CustomerID, inv.Amount,
        inv.Currency, inv.Status, inv.IssueDate, inv.DueDate, inv.Notes,
    ).Scan(&inv.ID, &inv.CreatedAt, &inv.UpdatedAt)
}

func (r *InvoiceRepository) CreateInvoiceItem(ctx context.Context, tx *sql.Tx, invoiceID int64, item struct {
    Description string
    Quantity    int
    Rate        float64
}) error {
    query := `
        INSERT INTO invoice_items (invoice_id, description, quantity, rate, amount)
        VALUES ($1, $2, $3, $4, $5)`

    amount := float64(item.Quantity) * item.Rate
    _, err := tx.ExecContext(ctx, query,
        invoiceID, item.Description, item.Quantity, item.Rate, amount)
    return err
}

func (r *InvoiceRepository) GetByID(ctx context.Context, id int64) (*Invoice, error) {
    query := `
        SELECT id, reference_number, sender_id, customer_id, amount, currency,
               status, issue_date, due_date, notes, created_at, updated_at
        FROM invoices
        WHERE id = $1`

    inv := &Invoice{}
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

func (r *InvoiceRepository) UpdateInvoiceTx(ctx context.Context, tx *sql.Tx, inv *Invoice) error {
    query := `
        UPDATE invoices
        SET status = $1, updated_at = $2
        WHERE id = $3`

    _, err := tx.ExecContext(ctx, query, inv.Status, inv.UpdatedAt, inv.ID)
    return err
}

func (r *InvoiceRepository) GetDashboardStats(ctx context.Context, userID int64) (map[string]struct {
    Count int
    Total float64
}, error) {
    query := `
        SELECT status, COUNT(*) as count, COALESCE(SUM(amount), 0) as total_amount
        FROM invoices
        WHERE sender_id = $1
        GROUP BY status`

    rows, err := r.db.QueryContext(ctx, query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    stats := make(map[string]struct {
        Count int
        Total float64
    })

    for rows.Next() {
        var status string
        var count int
        var total float64
        if err := rows.Scan(&status, &count, &total); err != nil {
            return nil, err
        }
        stats[status] = struct {
            Count int
            Total float64
        }{count, total}
    }

    return stats, nil
}

func (r *InvoiceRepository) LogActivityTx(ctx context.Context, tx *sql.Tx, activity *Activity) error {
    query := `
        INSERT INTO activities (user_id, invoice_id, action, details)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at`

    return tx.QueryRowContext(ctx, query,
        activity.UserID, activity.InvoiceID, activity.Action, activity.Details,
    ).Scan(&activity.ID, &activity.CreatedAt)
}
