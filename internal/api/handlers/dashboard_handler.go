package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/Tabintel/invoice-system/internal/service"
)

type DashboardHandler struct {
    service *service.DashboardService
}

func NewDashboardHandler(service *service.DashboardService) *DashboardHandler {
    return &DashboardHandler{service: service}
}

func (h *DashboardHandler) GetDashboard(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    userID := r.Context().Value("user_id").(int64)

    dashboard, err := h.service.GetDashboard(r.Context(), userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(dashboard)
}
