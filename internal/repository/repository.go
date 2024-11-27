package repository

import (
    "context"
    "database/sql"
   // "time"
)

type Repository struct {
    db *sql.DB
}

// Core methods
func NewRepository(db *sql.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) BeginTx(ctx context.Context) (*sql.Tx, error) {
    return r.db.BeginTx(ctx, nil)
}

// Implement GetDashboardStats
func (r *Repository) GetDashboardStats(ctx context.Context, userID int64) (map[string]interface{}, error) {
    // Dummy implementation for demonstration
    stats := map[string]interface{}{
        "total_invoices": 10,
        "total_amount":   1000.0,
    }
    return stats, nil
}

// User methods
func (r *Repository) GetByEmail(ctx context.Context, email string) (*User, error) {
    query := `SELECT id, name, email, password, company_name FROM users WHERE email = $1`
    user := &User{}
    err := r.db.QueryRowContext(ctx, query, email).Scan(
        &user.ID, &user.Name, &user.Email, &user.Password, &user.CompanyName)
    return user, err
}

func (r *Repository) Create(ctx context.Context, user *User) error {
    query := `INSERT INTO users (name, email, password, company_name) 
              VALUES ($1, $2, $3, $4) RETURNING id`
    return r.db.QueryRowContext(ctx, query, 
        user.Name, user.Email, user.Password, user.CompanyName).Scan(&user.ID)
}

// Invoice methods
func (r *Repository) CreateInvoice(ctx context.Context, tx *sql.Tx, inv *Invoice) error {
    query := `INSERT INTO invoices (reference_number, sender_id, customer_id, amount, 
              currency, status, issue_date, due_date, notes)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
              RETURNING id, created_at, updated_at`
    return tx.QueryRowContext(ctx, query,
        inv.ReferenceNumber, inv.SenderID, inv.CustomerID, inv.Amount,
        inv.Currency, inv.Status, inv.IssueDate, inv.DueDate, inv.Notes,
    ).Scan(&inv.ID, &inv.CreatedAt, &inv.UpdatedAt)
}

func (r *Repository) GetRecentInvoices(ctx context.Context, userID int64, limit int) ([]Invoice, error) {
    query := `SELECT * FROM invoices WHERE sender_id = $1 ORDER BY created_at DESC LIMIT $2`
    rows, err := r.db.QueryContext(ctx, query, userID, limit)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var invoices []Invoice
    for rows.Next() {
        var inv Invoice
        if err := rows.Scan(&inv.ID, &inv.ReferenceNumber, &inv.SenderID, &inv.CustomerID,
            &inv.Amount, &inv.Currency, &inv.Status, &inv.IssueDate, &inv.DueDate,
            &inv.Notes, &inv.CreatedAt, &inv.UpdatedAt); err != nil {
            return nil, err
        }
        invoices = append(invoices, inv)
    }
    return invoices, nil
}

// Activity methods
func (r *Repository) LogActivity(ctx context.Context, tx *sql.Tx, activity *Activity) error {
    query := `INSERT INTO activities (user_id, invoice_id, action, details)
              VALUES ($1, $2, $3, $4) RETURNING id, created_at`
    return tx.QueryRowContext(ctx, query,
        activity.UserID, activity.InvoiceID, activity.Action, activity.Details,
    ).Scan(&activity.ID, &activity.CreatedAt)
}

func (r *Repository) GetRecentActivities(ctx context.Context, userID int64, limit int) ([]Activity, error) {
    query := `SELECT a.*, i.reference_number FROM activities a
              JOIN invoices i ON a.invoice_id = i.id
              WHERE a.user_id = $1 ORDER BY a.created_at DESC LIMIT $2`
    rows, err := r.db.QueryContext(ctx, query, userID, limit)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var activities []Activity
    for rows.Next() {
        var act Activity
        if err := rows.Scan(&act.ID, &act.UserID, &act.InvoiceID, &act.Action,
            &act.Details, &act.CreatedAt, &act.InvoiceNumber); err != nil {
            return nil, err
        }
        activities = append(activities, act)
    }
    return activities, nil
}
