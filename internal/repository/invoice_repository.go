package repository

import (
    "context"
    "github.com/Tabintel/invoice-system/internal/ent"
    "github.com/Tabintel/invoice-system/internal/ent/invoice"
)

type InvoiceRepository struct {
    client *ent.Client
}

func NewInvoiceRepository(client *ent.Client) *InvoiceRepository {
    return &InvoiceRepository{client: client}
}

func (r *InvoiceRepository) Create(ctx context.Context, inv *ent.Invoice) (*ent.Invoice, error) {
    return r.client.Invoice.
        Create().
        SetReferenceNumber(inv.ReferenceNumber).
        SetAmount(inv.Amount).
        SetStatus(inv.Status).
        SetDueDate(inv.DueDate).
        Save(ctx)
}

func (r *InvoiceRepository) GetByID(ctx context.Context, id int) (*ent.Invoice, error) {
    return r.client.Invoice.Query().Where(invoice.ID(id)).Only(ctx)
}
