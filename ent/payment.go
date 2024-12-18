// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Tabintel/invoice-system/ent/payment"
)

// Payment is the model entity for the Payment schema.
type Payment struct {
	config
	// ID of the ent.
	ID               int `json:"id,omitempty"`
	invoice_payments *int
	selectValues     sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Payment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case payment.FieldID:
			values[i] = new(sql.NullInt64)
		case payment.ForeignKeys[0]: // invoice_payments
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Payment fields.
func (pa *Payment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case payment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = int(value.Int64)
		case payment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field invoice_payments", value)
			} else if value.Valid {
				pa.invoice_payments = new(int)
				*pa.invoice_payments = int(value.Int64)
			}
		default:
			pa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Payment.
// This includes values selected through modifiers, order, etc.
func (pa *Payment) Value(name string) (ent.Value, error) {
	return pa.selectValues.Get(name)
}

// Update returns a builder for updating this Payment.
// Note that you need to call Payment.Unwrap() before calling this method if this Payment
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Payment) Update() *PaymentUpdateOne {
	return NewPaymentClient(pa.config).UpdateOne(pa)
}

// Unwrap unwraps the Payment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Payment) Unwrap() *Payment {
	_tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Payment is not a transactional entity")
	}
	pa.config.driver = _tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Payment) String() string {
	var builder strings.Builder
	builder.WriteString("Payment(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Payments is a parsable slice of Payment.
type Payments []*Payment
