package service

import (
    "bytes"
    "fmt"
    "time"
    "github.com/jung-kurt/gofpdf"
)

type PDFService struct {
    invoiceRepo *repository.InvoiceRepository
}

func NewPDFService(invoiceRepo *repository.InvoiceRepository) *PDFService {
    return &PDFService{invoiceRepo: invoiceRepo}
}

func (s *PDFService) GenerateInvoicePDF(invoice *repository.Invoice, sender, customer *repository.User) (*bytes.Buffer, error) {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    
    // Header
    pdf.SetFont("Arial", "B", 16)
    pdf.Cell(190, 10, "INVOICE")
    
    // Invoice Details
    pdf.SetFont("Arial", "", 12)
    pdf.Ln(20)
    pdf.Cell(190, 10, fmt.Sprintf("Invoice #: %s", invoice.ReferenceNumber))
    pdf.Ln(10)
    pdf.Cell(190, 10, fmt.Sprintf("Date: %s", invoice.IssueDate.Format("02/01/2006")))
    pdf.Ln(10)
    pdf.Cell(190, 10, fmt.Sprintf("Due Date: %s", invoice.DueDate.Format("02/01/2006")))
    
    // Company Details
    pdf.Ln(20)
    pdf.SetFont("Arial", "B", 12)
    pdf.Cell(95, 10, "From:")
    pdf.Cell(95, 10, "Bill To:")
    pdf.Ln(10)
    pdf.SetFont("Arial", "", 12)
    pdf.Cell(95, 10, sender.CompanyName)
    pdf.Cell(95, 10, customer.CompanyName)
    
    // Items Table
    pdf.Ln(20)
    s.drawItemsTable(pdf, invoice.Items)
    
    // Totals
    pdf.Ln(10)
    pdf.Cell(150, 10, "Total:")
    pdf.Cell(40, 10, fmt.Sprintf("%s %.2f", invoice.Currency, invoice.Amount))
    
    var buf bytes.Buffer
    err := pdf.Output(&buf)
    if err != nil {
        return nil, err
    }
    
    return &buf, nil
}

func (s *PDFService) drawItemsTable(pdf *gofpdf.Fpdf, items []repository.InvoiceItem) {
    // Table headers
    pdf.SetFont("Arial", "B", 12)
    pdf.Cell(80, 10, "Description")
    pdf.Cell(30, 10, "Quantity")
    pdf.Cell(40, 10, "Rate")
    pdf.Cell(40, 10, "Amount")
    pdf.Ln(10)
    
    // Table content
    pdf.SetFont("Arial", "", 12)
    for _, item := range items {
        pdf.Cell(80, 10, item.Description)
        pdf.Cell(30, 10, fmt.Sprintf("%d", item.Quantity))
        pdf.Cell(40, 10, fmt.Sprintf("%.2f", item.Rate))
        pdf.Cell(40, 10, fmt.Sprintf("%.2f", item.Amount))
        pdf.Ln(10)
    }
}
