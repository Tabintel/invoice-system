package services

import (
    "bytes"
    "fmt"
    "github.com/jung-kurt/gofpdf"
    "github.com/Tabintel/invoice-system/ent"
)

type PDFService struct {
    invoiceService *InvoiceService
}

func NewPDFService(invoiceService *InvoiceService) *PDFService {
    return &PDFService{
        invoiceService: invoiceService,
    }
}

func (s *PDFService) GenerateInvoicePDF(invoice *ent.Invoice) ([]byte, error) {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    
    // Add company logo
    pdf.SetFont("Arial", "B", 16)
    pdf.Cell(40, 10, "INVOICE")
    
    // Add invoice details
    pdf.SetFont("Arial", "", 12)
    pdf.Ln(10)
    pdf.Cell(40, 10, fmt.Sprintf("Invoice #: %s", invoice.ReferenceNumber))
    pdf.Ln(10)
    pdf.Cell(40, 10, fmt.Sprintf("Date: %s", invoice.IssueDate.Format("2006-01-02")))
    
    var buf bytes.Buffer
    err := pdf.Output(&buf)
    if err != nil {
        return nil, err
    }
    
    return buf.Bytes(), nil
}