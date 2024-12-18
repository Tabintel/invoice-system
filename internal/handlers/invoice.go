package handlers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "log"
    "strconv"
    "github.com/go-chi/chi/v5"
    "github.com/Tabintel/invoice-system/ent"
    "github.com/Tabintel/invoice-system/internal/services"
)
type InvoiceHandler struct {
    service *services.InvoiceService
}

func NewInvoiceHandler(service *services.InvoiceService) *InvoiceHandler {
    return &InvoiceHandler{service: service}
}

// @Summary Create invoice
// @Description Create a new invoice
// @Tags invoices
// @Accept json
// @Produce json
// @Param invoice body services.CreateInvoiceInput true "Invoice Creation Input"
// @Success 200 {object} ent.Invoice
// @Router /invoices [post]
func (h *InvoiceHandler) Create() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var input services.CreateInvoiceInput
        if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
            log.Printf("Error decoding request: %v", err)
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }

        invoice, err := h.service.CreateInvoice(r.Context(), input)
        if err != nil {
            log.Printf("Error creating invoice: %v", err)
            http.Error(w, "Failed to create invoice", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "status": "success",
            "data":   invoice,
        })
    }
}

// @Summary List invoices
// @Description Get all invoices with optional filters
// @Tags invoices
// @Produce json
// @Success 200 {array} ent.Invoice
// @Router /invoices [get]
func (h *InvoiceHandler) List() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        input := services.ListInvoicesInput{
            Page:     1,
            PageSize: 10,
            Status:   r.URL.Query().Get("status"),
            SortBy:   r.URL.Query().Get("sort_by"),
        }
        
        invoices, err := h.service.ListInvoices(r.Context(), input)
        if err != nil {
            log.Printf("Error listing invoices: %v", err)
            http.Error(w, "Failed to list invoices", http.StatusInternalServerError)
            return
        }
        
        stats, err := h.service.GetStats(r.Context())
        if err != nil {
            log.Printf("Error getting invoice stats: %v", err)
            http.Error(w, "Failed to get invoice stats", http.StatusInternalServerError)
            return
        }
        
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "status": "success",
            "data": map[string]interface{}{
                "invoices": invoices,
                "stats":   stats,
            },
        })
    }
}

func (h *InvoiceHandler) UpdateStatus() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        idStr := chi.URLParam(r, "id")
        id, err := strconv.Atoi(idStr)
        if err != nil {
            http.Error(w, "Invalid invoice ID", http.StatusBadRequest)
            return
        }

        var input services.UpdateInvoiceStatusInput
        if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }

        invoice, err := h.service.UpdateStatus(r.Context(), id, input)
        if err != nil {
            log.Printf("Error updating invoice status: %v", err)
            http.Error(w, "Failed to update invoice status", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "status": "success",
            "data":   invoice,
        })
    }
}




func (h *InvoiceHandler) DownloadPDF() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id, _ := strconv.Atoi(chi.URLParam(r, "id"))
        
        invoice, err := h.service.GetInvoice(r.Context(), id)
        if err != nil {
            http.Error(w, "Invoice not found", http.StatusNotFound)
            return
        }
        
        pdfService := services.NewPDFService(h.service)
        pdfBytes, err := pdfService.GenerateInvoicePDF(invoice)
        if err != nil {
            http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
            return
        }
        
        w.Header().Set("Content-Type", "application/pdf")
        w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.pdf", invoice.ReferenceNumber))
        w.Write(pdfBytes)
    }
}

func (h *InvoiceHandler) GenerateShareableLink() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id, _ := strconv.Atoi(chi.URLParam(r, "id"))
        
        shareableLink, err := h.service.GenerateShareableLink(r.Context(), id)
        if err != nil {
            if ent.IsNotFound(err) {
                http.Error(w, "Invoice not found", http.StatusNotFound)
                return
            }
            log.Printf("Error generating shareable link: %v", err)
            http.Error(w, "Failed to generate link", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "status": "success",
            "data":   shareableLink,
        })
    }
}

func (h *InvoiceHandler) GetStats() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        stats, err := h.service.GetStats(r.Context())
        if err != nil {
            log.Printf("Error getting invoice stats: %v", err)
            http.Error(w, "Failed to get stats", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "status": "success",
            "data":   stats,
        })
    }
}