package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/Tabintel/invoice-system/internal/service"
)

type ActivityHandler struct {
    service *service.ActivityService
}

func NewActivityHandler(service *service.ActivityService) *ActivityHandler {
    return &ActivityHandler{service: service}
}

func (h *ActivityHandler) GetRecentActivities(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    userID := r.Context().Value("user_id").(int64)
    
    activities, err := h.service.GetRecentActivities(r.Context(), userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(activities)
}

func (h *ActivityHandler) GetInvoiceActivities(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    invoiceID := r.URL.Query().Get("invoice_id")
    if invoiceID == "" {
        http.Error(w, "Invoice ID is required", http.StatusBadRequest)
        return
    }

    userID := r.Context().Value("user_id").(int64)
    
    activities, err := h.service.GetInvoiceActivities(r.Context(), userID, invoiceID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(activities)
}
