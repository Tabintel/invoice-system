package repository

import (
    "context"
    "database/sql"
)

type Repository struct {
    db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
    return &Repository{db: db}
}

// Invoice methods
func (r *Repository) CreateInvoice(ctx context.Context, tx *sql.Tx, inv *Invoice) error {
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

func (r *Repository) GetInvoiceByID(ctx context.Context, id int64) (*Invoice, error) {
    query := `
        SELECT id, reference_number, sender_id, customer_id, amount, currency,
               status, issue_date, due_date, notes, created_at, updated_at
        FROM invoices WHERE id = $1`

    inv := &Invoice{}
    err := r.db.QueryRowContext(ctx, query, id).Scan(
        &inv.ID, &inv.ReferenceNumber, &inv.SenderID, &inv.CustomerID,
        &inv.Amount, &inv.Currency, &inv.Status, &inv.IssueDate,
        &inv.DueDate, &inv.Notes, &inv.CreatedAt, &inv.UpdatedAt,
    )
    return inv, err
}

func (r *Repository) UpdateInvoiceStatus(ctx context.Context, tx *sql.Tx, id int64, status string) error {
    query := `UPDATE invoices SET status = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2`
    _, err := tx.ExecContext(ctx, query, status, id)
    return err
}

// Activity methods
func (r *Repository) LogActivity(ctx context.Context, tx *sql.Tx, activity *Activity) error {
    query := `
        INSERT INTO activities (user_id, invoice_id, action, details)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at`

    return tx.QueryRowContext(ctx, query,
        activity.UserID, activity.InvoiceID, activity.Action, activity.Details,
    ).Scan(&activity.ID, &activity.CreatedAt)
}

func (r *Repository) GetRecentActivities(ctx context.Context, userID int64, limit int) ([]Activity, error) {
    query := `
        SELECT a.id, a.user_id, a.invoice_id, a.action, a.details, a.created_at,
               i.reference_number as invoice_number
        FROM activities a
        JOIN invoices i ON a.invoice_id = i.id
        WHERE a.user_id = $1
        ORDER BY a.created_at DESC
        LIMIT $2`

    rows, err := r.db.QueryContext(ctx, query, userID, limit)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var activities []Activity
    for rows.Next() {
        var act Activity
        err := rows.Scan(
            &act.ID, &act.UserID, &act.InvoiceID, &act.Action,
            &act.Details, &act.CreatedAt, &act.InvoiceNumber,
        )
        if err != nil {
            return nil, err
        }
        activities = append(activities, act)
    }
    return activities, nil
}

// User methods
func (r *Repository) CreateUser(ctx context.Context, user *User) error {
    query := `
        INSERT INTO users (name, email, password, company_name)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at`

    return r.db.QueryRowContext(ctx, query,
        user.Name, user.Email, user.Password, user.CompanyName,
    ).Scan(&user.ID, &user.CreatedAt)
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
    query := `SELECT id, name, email, password, company_name, created_at FROM users WHERE email = $1`
    
    user := &User{}
    err := r.db.QueryRowContext(ctx, query, email).Scan(
        &user.ID, &user.Name, &user.Email, &user.Password,
        &user.CompanyName, &user.CreatedAt,
    )
    return user, err
}

func (r *Repository) BeginTx(ctx context.Context) (*sql.Tx, error) {
    return r.db.BeginTx(ctx, nil)
}
