package service

import (
    "bytes"
    "html/template"
    "net/smtp"
    "github.com/Tabintel/invoice-system/internal/repository"
)

type NotificationService struct {
    emailFrom string
    emailPass string
    smtpHost  string
    smtpPort  string
}

func NewNotificationService(emailFrom, emailPass, smtpHost, smtpPort string) *NotificationService {
    return &NotificationService{
        emailFrom: emailFrom,
        emailPass: emailPass,
        smtpHost:  smtpHost,
        smtpPort:  smtpPort,
    }
}

func (s *NotificationService) SendInvoiceStatusUpdate(invoice *repository.Invoice, status string) error {
    template := `
        Invoice #{{.ReferenceNumber}} status has been updated to {{.Status}}
        Amount: {{.Currency}} {{.Amount}}
        Due Date: {{.DueDate}}
    `
    
    var body bytes.Buffer
    t := template.Must(template.New("invoice").Parse(template))
    t.Execute(&body, invoice)
    
    return s.sendEmail(invoice.CustomerEmail, "Invoice Status Update", body.String())
}
