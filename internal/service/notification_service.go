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
    tmpl := `
        Invoice #{{.ReferenceNumber}} status has been updated to {{.Status}}
        Amount: {{.Currency}} {{.Amount}}
        Due Date: {{.DueDate}}
    `
    
    var body bytes.Buffer
    t := template.Must(template.New("invoice").Parse(tmpl))
    t.Execute(&body, invoice)
    
    return s.sendEmail(invoice.CustomerEmail, "Invoice Status Update", body.String())
}

func (s *NotificationService) sendEmail(to, subject, body string) error {
    msg := "From: " + s.emailFrom + "\n" +
        "To: " + to + "\n" +
        "Subject: " + subject + "\n\n" +
        body

    auth := smtp.PlainAuth("", s.emailFrom, s.emailPass, s.smtpHost)
    return smtp.SendMail(s.smtpHost+":"+s.smtpPort, auth, s.emailFrom, []string{to}, []byte(msg))
}
