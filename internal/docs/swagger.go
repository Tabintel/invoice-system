package docs

import (
    "github.com/go-openapi/spec"
)

func GetSwaggerSpec() *spec.Swagger {
    swagger := &spec.Swagger{
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
                Paths: map[string]spec.PathItem{
                    "/api/invoices": {
                        PathItemProps: spec.PathItemProps{
                            Post: &spec.Operation{
                                OperationProps: spec.OperationProps{
                                    Description: "Create new invoice",
                                    Tags:        []string{"invoices"},
                                    Parameters: []spec.Parameter{
                                        {
                                            ParamProps: spec.ParamProps{
                                                Name:     "invoice",
                                                In:       "body",
                                                Required: true,
                                                Schema: &spec.Schema{
                                                    SchemaProps: spec.SchemaProps{
                                                        Ref: spec.MustCreateRef("#/definitions/CreateInvoiceInput"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                            Get: &spec.Operation{
                                OperationProps: spec.OperationProps{
                                    Description: "List all invoices",
                                    Tags:        []string{"invoices"},
                                },
                            },
                        },
                    },
                },
            },
            "/api/invoices/{id}/status": {
                PathItemProps: spec.PathItemProps{
                    Put: &spec.Operation{
                        OperationProps: spec.OperationProps{
                            Description: "Update invoice status",
                            Tags:        []string{"invoices"},
                            Parameters: []spec.Parameter{
                                {
                                    ParamProps: spec.ParamProps{
                                        Name:     "id",
                                        In:       "path",
                                        Required: true,
                                        Type:     "integer",
                                    },
                                },
                                {
                                    ParamProps: spec.ParamProps{
                                        Name:     "status",
                                        In:       "body",
                                        Required: true,
                                        Schema: &spec.Schema{
                                            SchemaProps: spec.SchemaProps{
                                                Ref: spec.MustCreateRef("#/definitions/UpdateInvoiceStatusInput"),
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
            },
        },
    }
    
    swagger.Definitions = spec.Definitions{
        "CreateInvoiceInput": spec.Schema{
            SchemaProps: spec.SchemaProps{
                Type: []string{"object"},
                Properties: map[string]spec.Schema{
                    "customer_id": {
                        SchemaProps: spec.SchemaProps{
                            Type: []string{"integer"},
                        },
                    },
                    "due_date": {
                        SchemaProps: spec.SchemaProps{
                            Type: []string{"string"},
                            Format: "date-time",
                        },
                    },
                    "currency": {
                        SchemaProps: spec.SchemaProps{
                            Type: []string{"string"},
                        },
                    },
                    "items": {
                        SchemaProps: spec.SchemaProps{
                            Type: []string{"array"},
                            Items: &spec.SchemaOrArray{
                                Schema: &spec.Schema{
                                    SchemaProps: spec.SchemaProps{
                                        Ref: spec.MustCreateRef("#/definitions/InvoiceItemInput"),
                                    },
                                },
                            },
                        },
                    },
                },
            },
        },
        "InvoiceItemInput": spec.Schema{
            SchemaProps: spec.SchemaProps{
                Properties: map[string]spec.Schema{
                    "description": {Type: []string{"string"}},
                    "quantity": {Type: []string{"integer"}},
                    "rate": {Type: []string{"number"}},
                },
            },
        },
    }
    
    return swagger
}