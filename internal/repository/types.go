package repository

import "time"

type Invoice struct {
    ID              int64     `json:"id"`
    ReferenceNumber string    `json:"reference_number"`
    SenderID        int64     `json:"sender_id"`
    CustomerID      int64     `json:"customer_id"`
    CustomerEmail   string    `json:"customer_email"`
    Amount          float64   `json:"amount"`
    Currency        string    `json:"currency"`
    Status          string    `json:"status"`
    IssueDate       time.Time `json:"issue_date"`
    DueDate         time.Time `json:"due_date"`
    Notes           string    `json:"notes"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
    Items           []InvoiceItem `json:"items,omitempty"`
}

type InvoiceItem struct {
    ID          int64   `json:"id"`
    InvoiceID   int64   `json:"invoice_id"`
    Description string  `json:"description"`
    Quantity    int     `json:"quantity"`
    Rate        float64 `json:"rate"`
    Amount      float64 `json:"amount"`
}
