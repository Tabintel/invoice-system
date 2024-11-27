package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/Tabintel/invoice-system/internal/service"
)

type InvoiceHandler struct {
    service *service.InvoiceService
}

func NewInvoiceHandler(service *service.InvoiceService) *InvoiceHandler {
    return &InvoiceHandler{service: service}
}

func (h *InvoiceHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value("user_id").(int64)

    var req struct {
        CustomerDetails struct {
            Name    string `json:"name"`
            Email   string `json:"email"`
            Phone   string `json:"phone"`
        } `json:"customer_details"`
        Items []struct {
            Description string  `json:"description"`
            Quantity    int     `json:"quantity"`
            Rate       float64 `json:"rate"`
        } `json:"items"`
        Currency    string    `json:"currency"`
        IssueDate   time.Time `json:"issue_date"`
        DueDate     time.Time `json:"due_date"`
        Notes       string    `json:"notes"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    invoice, err := h.service.CreateInvoice(r.Context(), &service.CreateInvoiceRequest{
        SenderID:        userID,
        CustomerDetails: req.CustomerDetails,
        Items:          req.Items,
        Currency:       req.Currency,
        IssueDate:      req.IssueDate,
        DueDate:        req.DueDate,
        Notes:          req.Notes,
    })

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(invoice)
}