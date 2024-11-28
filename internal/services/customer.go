package services

import (
    "context"
    "github.com/Tabintel/invoice-system/ent"
    "github.com/Tabintel/invoice-system/ent/customer"
)

type CustomerService struct {
    db *ent.Client
}

type CreateCustomerInput struct {
    Name    string `json:"name"`
    Email   string `json:"email"`
    Phone   string `json:"phone"`
    Address string `json:"address"`
}

func NewCustomerService(db *ent.Client) *CustomerService {
    return &CustomerService{db: db}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, input CreateCustomerInput) (*ent.Customer, error) {
    return s.db.Customer.Create().
        SetName(input.Name).
        SetEmail(input.Email).
        SetPhone(input.Phone).
        SetAddress(input.Address).
        Save(ctx)
}

func (s *CustomerService) ListCustomers(ctx context.Context) ([]*ent.Customer, error) {
    return s.db.Customer.Query().
        Order(ent.Desc(customer.FieldCreatedAt)).
        All(ctx)
}
