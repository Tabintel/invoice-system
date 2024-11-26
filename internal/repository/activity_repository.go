package repository

import (
    "context"
    "github.com/Tabintel/invoice-system/internal/ent"
)

type ActivityRepository struct {
    client *ent.Client
}

func (r *ActivityRepository) LogActivity(ctx context.Context, activity *ent.ActivityLog) error {
    _, err := r.client.ActivityLog.Create().
        SetActionType(activity.ActionType).
        SetDescription(activity.Description).
        SetPerformedBy(activity.PerformedBy).
        Save(ctx)
    return err
}

func (r *ActivityRepository) GetRecentActivities(ctx context.Context, limit int) ([]*ent.ActivityLog, error) {
    return r.client.ActivityLog.Query().
        Order(ent.Desc("timestamp")).
        Limit(limit).
        All(ctx)
}
