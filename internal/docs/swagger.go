package docs

import (
    "github.com/go-openapi/spec"
   // "github.com/go-openapi/strfmt"
)

func GetSwaggerSpec() *spec.Swagger {
    paths := make(map[string]spec.PathItem)
    paths["/api/invoices"] = spec.PathItem{
        PathItemProps: spec.PathItemProps{
            Post: &spec.Operation{
                OperationProps: spec.OperationProps{
                    Description: "Create new invoice",
                    Tags:        []string{"invoices"},
                },
            },
            Get: &spec.Operation{
                OperationProps: spec.OperationProps{
                    Description: "List all invoices",
                    Tags:        []string{"invoices"},
                },
            },
        },
    }
    
    paths["/api/invoices/{id}/status"] = spec.PathItem{
        PathItemProps: spec.PathItemProps{
            Put: &spec.Operation{
                OperationProps: spec.OperationProps{
                    Description: "Update invoice status",
                    Tags:        []string{"invoices"},
                },
            },
        },
    }

    definitions := spec.Definitions{
        "CreateInvoiceInput": {
            SchemaProps: spec.SchemaProps{
                Properties: map[string]spec.Schema{
                    "customer_id": {
                        SchemaProps: spec.SchemaProps{
                            Type: spec.StringOrArray{"integer"},
                        },
                    },
                    "due_date": {
                        SchemaProps: spec.SchemaProps{
                            Type:   spec.StringOrArray{"string"},
                            Format: "date-time",
                        },
                    },
                },
            },
        },
    }

    return &spec.Swagger{
        SwaggerProps: spec.SwaggerProps{
            Swagger: "2.0",
            Info: &spec.Info{
                InfoProps: spec.InfoProps{
                    Title:       "Invoice System API",
                    Description: "Modern invoice management system API",
                    Version:     "1.0.0",
                },
            },
            Paths: &spec.Paths{
                Paths: paths,
            },
            Definitions: definitions,
        },
    }
}