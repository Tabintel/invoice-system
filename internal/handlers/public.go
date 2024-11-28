package handlers

import (
    //"encoding/json"
    "net/http"
    "html/template"
    //"time"
    "github.com/go-chi/chi/v5"
    "github.com/Tabintel/invoice-system/internal/services"
)

type PublicHandler struct {
    service *services.InvoiceService
}

func NewPublicHandler(service *services.InvoiceService) *PublicHandler {
    return &PublicHandler{service: service}
}

func (h *PublicHandler) ViewPublicInvoice() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := chi.URLParam(r, "token")
        
        invoice, err := h.service.GetInvoiceByShareToken(r.Context(), token)
        if err != nil {
            http.Error(w, "Invoice not found", http.StatusNotFound)
            return
        }
        
        // Set HTML content type
        w.Header().Set("Content-Type", "text/html")
        
        // Render beautiful invoice template
        tmpl := template.Must(template.ParseFiles("internal/templates/invoice.html"))
        tmpl.Execute(w, map[string]interface{}{
            "Invoice": invoice,
            "CompanyName": "Your Company",
            "CompanyLogo": "/static/logo.png",
        })
    }
}