package repository

import (
    "context"
    "database/sql"
    "time"
)

type Repository struct {
    db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
    return &Repository{db: db}
}

// Activity Methods
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
