package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/Tabintel/invoice-system/internal/service"
)

type ShareHandler struct {
    shareService   *service.ShareService
    invoiceService *service.InvoiceService
}

func NewShareHandler(shareService *service.ShareService, invoiceService *service.InvoiceService) *ShareHandler {
    return &ShareHandler{
        shareService:   shareService,
        invoiceService: invoiceService,
    }
}

func (h *ShareHandler) GenerateShareableLink(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req struct {
        InvoiceID   int64 `json:"invoice_id"`
        ExpiryHours int   `json:"expiry_hours"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    userID := r.Context().Value("user_id").(int64)

    // Verify invoice ownership
    if err := h.invoiceService.VerifyOwnership(r.Context(), req.InvoiceID, userID); err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    link, err := h.shareService.GenerateShareableLink(req.InvoiceID, req.ExpiryHours)
    if err != nil {
        http.Error(w, "Failed to generate link", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(link)
}

func (h *ShareHandler) GetSharedInvoice(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    token := r.URL.Query().Get("token")
    if token == "" {
        http.Error(w, "Token is required", http.StatusBadRequest)
        return
    }

    invoice, err := h.shareService.GetInvoiceByShareToken(r.Context(), token)
    if err != nil {
        http.Error(w, "Invalid or expired link", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(invoice)
}
