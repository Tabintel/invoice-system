package handlers

import (
    "fmt"
    "net/http"
    "strconv"
    "github.com/Tabintel/invoice-system/internal/service"
)

type PDFHandler struct {
    pdfService     *service.PDFService
    invoiceService *service.InvoiceService
}

func NewPDFHandler(pdfService *service.PDFService, invoiceService *service.InvoiceService) *PDFHandler {
    return &PDFHandler{
        pdfService:     pdfService,
        invoiceService: invoiceService,
    }
}

func (h *PDFHandler) GenerateInvoicePDF(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    invoiceID := r.URL.Query().Get("id")
    if invoiceID == "" {
        http.Error(w, "Invoice ID is required", http.StatusBadRequest)
        return
    }

    id, err := strconv.ParseInt(invoiceID, 10, 64)
    if err != nil {
        http.Error(w, "Invalid invoice ID", http.StatusBadRequest)
        return
    }

    userID := r.Context().Value("user_id").(int64)

    // Get invoice with items
    invoice, err := h.invoiceService.GetInvoiceWithDetails(r.Context(), id, userID)
    if err != nil {
        http.Error(w, "Failed to get invoice", http.StatusInternalServerError)
        return
    }

    // Generate PDF
    pdfBuffer, err := h.pdfService.GenerateInvoicePDF(invoice)
    if err != nil {
        http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
        return
    }

    // Set response headers
    w.Header().Set("Content-Type", "application/pdf")
    w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=invoice-%s.pdf", invoice.ReferenceNumber))
    w.Header().Set("Content-Length", strconv.Itoa(pdfBuffer.Len()))

    // Write PDF to response
    if _, err := pdfBuffer.WriteTo(w); err != nil {
        http.Error(w, "Failed to send PDF", http.StatusInternalServerError)
        return
    }
}
