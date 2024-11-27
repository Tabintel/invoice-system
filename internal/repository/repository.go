package repository

import (
    "context"
    "database/sql"
    //"time"
)

type Repository struct {
    db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) CreateOrGetCustomer(ctx context.Context, details struct {
    Name  string
    Email string
    Phone string
}) (int64, error) {
    var customerID int64
    query := `INSERT INTO users (name, email, phone, role) 
              VALUES ($1, $2, $3, 'customer') 
              ON CONFLICT (email) DO UPDATE SET name = $1, phone = $3 
              RETURNING id`
    err := r.db.QueryRowContext(ctx, query, details.Name, details.Email, details.Phone).Scan(&customerID)
    return customerID, err
}

func (r *Repository) CreateTx(ctx context.Context, tx *sql.Tx, invoice *Invoice) error {
    query := `INSERT INTO invoices (reference_number, sender_id, customer_id, amount, 
              currency, status, issue_date, due_date, notes)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
              RETURNING id, created_at, updated_at`
    return tx.QueryRowContext(ctx, query,
        invoice.ReferenceNumber, invoice.SenderID, invoice.CustomerID, invoice.Amount,
        invoice.Currency, invoice.Status, invoice.IssueDate, invoice.DueDate, invoice.Notes,
    ).Scan(&invoice.ID, &invoice.CreatedAt, &invoice.UpdatedAt)
}

func (r *Repository) CreateInvoiceItem(ctx context.Context, tx *sql.Tx, invoiceID int64, item struct {
    Description string
    Quantity    int
    Rate        float64
}) error {
    query := `INSERT INTO invoice_items (invoice_id, description, quantity, rate, amount)
              VALUES ($1, $2, $3, $4, $5)`
    amount := float64(item.Quantity) * item.Rate
    _, err := tx.ExecContext(ctx, query, invoiceID, item.Description, item.Quantity, item.Rate, amount)
    return err
}

func (r *Repository) GetByID(ctx context.Context, id int64) (*Invoice, error) {
    query := `SELECT i.*, u.email as customer_email FROM invoices i 
              JOIN users u ON i.customer_id = u.id WHERE i.id = $1`
    inv := &Invoice{}
    err := r.db.QueryRowContext(ctx, query, id).Scan(
        &inv.ID, &inv.ReferenceNumber, &inv.SenderID, &inv.CustomerID,
        &inv.Amount, &inv.Currency, &inv.Status, &inv.IssueDate,
        &inv.DueDate, &inv.Notes, &inv.CreatedAt, &inv.UpdatedAt,
        &inv.CustomerEmail,
    )
    return inv, err
}

func (r *Repository) UpdateInvoiceTx(ctx context.Context, tx *sql.Tx, invoice *Invoice) error {
    query := `UPDATE invoices SET status = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2`
    _, err := tx.ExecContext(ctx, query, invoice.Status, invoice.ID)
    return err
}

func (r *Repository) LogActivityTx(ctx context.Context, tx *sql.Tx, activity *Activity) error {
    query := `INSERT INTO activities (user_id, invoice_id, action, details)
              VALUES ($1, $2, $3, $4) RETURNING id, created_at`
    return tx.QueryRowContext(ctx, query,
        activity.UserID, activity.InvoiceID, activity.Action, activity.Details,
    ).Scan(&activity.ID, &activity.CreatedAt)
}

func (r *Repository) SaveShareableLink(ctx context.Context, invoiceID int64, link string) error {
    query := `UPDATE invoices SET shareable_link = $1 WHERE id = $2`
    _, err := r.db.ExecContext(ctx, query, link, invoiceID)
    return err
}

func (r *Repository) BeginTx(ctx context.Context) (*sql.Tx, error) {
    return r.db.BeginTx(ctx, nil)
}

func (r *Repository) GetDashboardStats(ctx context.Context, userID int64) (map[string]interface{}, error) {
    stats := map[string]interface{}{
        "total_invoices": 10,
        "total_amount":   1000.0,
    }
    return stats, nil
}

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
