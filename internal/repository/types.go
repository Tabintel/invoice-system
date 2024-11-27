package repository

import (
    "time"
)

type User struct {
    ID          int64     `json:"id"`
    Name        string    `json:"name"`
    Email       string    `json:"email"`
    Password    string    `json:"-"`
    CompanyName string    `json:"company_name"`
    CreatedAt   time.Time `json:"created_at"`
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

type InvoiceInfo struct {
    ID              int64     `json:"id"`
    ReferenceNumber string    `json:"reference_number"`
    Amount          float64   `json:"amount"`
    Status          string    `json:"status"`
    DueDate         time.Time `json:"due_date"`
}

type ActivityInfo struct {
    ID            int64     `json:"id"`
    Action        string    `json:"action"`
    Details       string    `json:"details"`
    InvoiceNumber string    `json:"invoice_number"`
    Timestamp     time.Time `json:"timestamp"`
}