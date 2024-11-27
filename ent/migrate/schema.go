// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CustomersColumns holds the columns for the "customers" table.
	CustomersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// CustomersTable holds the schema information for the "customers" table.
	CustomersTable = &schema.Table{
		Name:       "customers",
		Columns:    CustomersColumns,
		PrimaryKey: []*schema.Column{CustomersColumns[0]},
	}
	// InvoicesColumns holds the columns for the "invoices" table.
	InvoicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "reference_number", Type: field.TypeString, Unique: true},
		{Name: "total_amount", Type: field.TypeFloat64},
		{Name: "status", Type: field.TypeString, Default: "draft"},
		{Name: "issue_date", Type: field.TypeTime},
		{Name: "due_date", Type: field.TypeTime},
		{Name: "currency", Type: field.TypeString, Default: "USD"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "invoice_customer", Type: field.TypeInt, Nullable: true},
		{Name: "user_invoices", Type: field.TypeInt, Nullable: true},
	}
	// InvoicesTable holds the schema information for the "invoices" table.
	InvoicesTable = &schema.Table{
		Name:       "invoices",
		Columns:    InvoicesColumns,
		PrimaryKey: []*schema.Column{InvoicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "invoices_customers_customer",
				Columns:    []*schema.Column{InvoicesColumns[8]},
				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "invoices_users_invoices",
				Columns:    []*schema.Column{InvoicesColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// InvoiceItemsColumns holds the columns for the "invoice_items" table.
	InvoiceItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "invoice_items", Type: field.TypeInt, Nullable: true},
	}
	// InvoiceItemsTable holds the schema information for the "invoice_items" table.
	InvoiceItemsTable = &schema.Table{
		Name:       "invoice_items",
		Columns:    InvoiceItemsColumns,
		PrimaryKey: []*schema.Column{InvoiceItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "invoice_items_invoices_items",
				Columns:    []*schema.Column{InvoiceItemsColumns[1]},
				RefColumns: []*schema.Column{InvoicesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PaymentsColumns holds the columns for the "payments" table.
	PaymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "invoice_payments", Type: field.TypeInt, Nullable: true},
	}
	// PaymentsTable holds the schema information for the "payments" table.
	PaymentsTable = &schema.Table{
		Name:       "payments",
		Columns:    PaymentsColumns,
		PrimaryKey: []*schema.Column{PaymentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "payments_invoices_payments",
				Columns:    []*schema.Column{PaymentsColumns[1]},
				RefColumns: []*schema.Column{InvoicesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "phone", Type: field.TypeString},
		{Name: "company_name", Type: field.TypeString},
		{Name: "role", Type: field.TypeString, Default: "user"},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CustomersTable,
		InvoicesTable,
		InvoiceItemsTable,
		PaymentsTable,
		UsersTable,
	}
)

func init() {
	InvoicesTable.ForeignKeys[0].RefTable = CustomersTable
	InvoicesTable.ForeignKeys[1].RefTable = UsersTable
	InvoiceItemsTable.ForeignKeys[0].RefTable = InvoicesTable
	PaymentsTable.ForeignKeys[0].RefTable = InvoicesTable
}
