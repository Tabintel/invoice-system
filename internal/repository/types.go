package repository

import (
    "time"
)

type User struct {
    ID          int64     `json:"id"`
    Name        string    `json:"name"`
    Email       string    `json:"email"`
    Password    string    `json:"-"`
    Phone       string    `json:"phone"`
    CompanyName string    `json:"company_name"`
    CompanyLogo string    `json:"company_logo"`
    Location    string    `json:"location"`
    Role        string    `json:"role"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
type Activity struct {
    ID            int64     `json:"id"`
    UserID        int64     `json:"user_id"`
    InvoiceID     int64     `json:"invoice_id"`
    Action        string    `json:"action"`
    Details       string    `json:"details"`
    CreatedAt     time.Time `json:"created_at"`
    InvoiceNumber string    `json:"invoice_number"`
}

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

type ActivityInfo struct {
    ID            int64     `json:"id"`
    Action        string    `json:"action"`
    Details       string    `json:"details"`
    InvoiceNumber string    `json:"invoice_number"`
    Timestamp     time.Time `json:"timestamp"`
}