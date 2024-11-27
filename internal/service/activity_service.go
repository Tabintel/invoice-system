package service

import (
    "context"
    "github.com/Tabintel/invoice-system/internal/repository"
)

type ActivityService struct {
    repo *repository.Repository
}

func NewActivityService(repo *repository.Repository) *ActivityService {
    return &ActivityService{repo: repo}
}

const (
    ActionCreated  = "created"
    ActionUpdated  = "updated"
    ActionSent     = "sent"
    ActionPaid     = "paid"
    ActionOverdue  = "marked_overdue"
    ActionShared   = "shared"
)

func (s *ActivityService) LogInvoiceActivity(ctx context.Context, tx *sql.Tx, userID, invoiceID int64, action string) error {
    details := s.getActivityDetails(action)
    
    activity := &repository.Activity{
        UserID:    userID,
        InvoiceID: invoiceID,
        Action:    action,
        Details:   details,
    }

    return s.repo.LogActivity(ctx, tx, activity)
}

func (s *ActivityService) GetRecentActivities(ctx context.Context, userID int64) ([]ActivityResponse, error) {
    activities, err := s.repo.GetRecentActivities(ctx, userID, 10)
    if err != nil {
        return nil, err
    }

    return s.formatActivities(activities), nil
}

type ActivityResponse struct {
    ID              int64     `json:"id"`
    Action          string    `json:"action"`
    Details         string    `json:"details"`
    InvoiceNumber   string    `json:"invoice_number"`
    Timestamp       time.Time `json:"timestamp"`
}

func (s *ActivityService) formatActivities(activities []repository.Activity) []ActivityResponse {
    response := make([]ActivityResponse, len(activities))
    for i, act := range activities {
        response[i] = ActivityResponse{
            ID:            act.ID,
            Action:        act.Action,
            Details:       act.Details,
            InvoiceNumber: act.InvoiceNumber,
            Timestamp:     act.CreatedAt,
        }
    }
    return response
}

func (s *ActivityService) getActivityDetails(action string) string {
    switch action {
    case ActionCreated:
        return "Invoice created"
    case ActionUpdated:
        return "Invoice updated"
    case ActionSent:
        return "Invoice sent to customer"
    case ActionPaid:
        return "Invoice marked as paid"
    case ActionOverdue:
        return "Invoice marked as overdue"
    case ActionShared:
        return "Invoice shared"
    default:
        return "Invoice activity recorded"
    }
}
