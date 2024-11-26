package handlers

import (
    "net/http"
    "github.com/gofiber/fiber/v2"
    "github.com/Tabintel/invoice-system/internal/service"
)

type InvoiceHandler struct {
    invoiceService *service.InvoiceService
}

func NewInvoiceHandler(invoiceService *service.InvoiceService) *InvoiceHandler {
    return &InvoiceHandler{
        invoiceService: invoiceService,
    }
}

// @Summary Create new invoice
// @Description Create a new invoice in the system
// @Tags invoices
// @Accept json
// @Produce json
// @Param invoice body CreateInvoiceRequest true "Invoice details"
// @Success 201 {object} ent.Invoice
// @Router /invoices [post]
func (h *InvoiceHandler) CreateInvoice(c *fiber.Ctx) error {
    var req CreateInvoiceRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    input := service.CreateInvoiceInput{
        Amount:     req.Amount,
        Currency:   req.Currency,
        DueDate:    req.DueDate,
        Notes:      req.Notes,
        CustomerID: req.CustomerID,
    }

    invoice, err := h.invoiceService.CreateInvoice(c.Context(), input)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.Status(http.StatusCreated).JSON(invoice)
}

type CreateInvoiceRequest struct {
    Amount     float64   `json:"amount"`
    Currency   string    `json:"currency"`
    DueDate    time.Time `json:"due_date"`
    Notes      string    `json:"notes"`
    CustomerID int       `json:"customer_id"`
}
