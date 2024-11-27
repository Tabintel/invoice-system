package repository

import (
    "context"
    "database/sql"
    "time"
)

type Activity struct {
    ID        int64     `json:"id"`
    UserID    int64     `json:"user_id"`
    InvoiceID int64     `json:"invoice_id"`
    Action    string    `json:"action"`
    Details   string    `json:"details"`
    CreatedAt time.Time `json:"created_at"`
}

func (r *InvoiceRepository) LogActivity(ctx context.Context, tx *sql.Tx, activity *Activity) error {
    query := `
        INSERT INTO activities (user_id, invoice_id, action, details)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at`

    return tx.QueryRowContext(ctx, query,
        activity.UserID,
        activity.InvoiceID,
        activity.Action,
        activity.Details,
    ).Scan(&activity.ID, &activity.CreatedAt)
}

func (r *InvoiceRepository) GetRecentActivities(ctx context.Context, userID int64, limit int) ([]Activity, error) {
    query := `
        SELECT a.id, a.user_id, a.invoice_id, a.action, a.details, a.created_at,
               i.reference_number, u.name as user_name
        FROM activities a
        JOIN invoices i ON a.invoice_id = i.id
        JOIN users u ON a.user_id = u.id
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
            &act.Details, &act.CreatedAt,
        )
        if err != nil {
            return nil, err
        }
        activities = append(activities, act)
    }

    return activities, nil
}
