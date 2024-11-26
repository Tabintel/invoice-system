package repository

import (
    "context"
    "log"
    
    "github.com/Tabintel/invoice-system/ent"
    _ "github.com/lib/pq"
)

func NewDatabase(connectionString string) *ent.Client {
    client, err := ent.Open("postgres", connectionString)
    if err != nil {
        log.Fatalf("failed opening connection to postgres: %v", err)
    }
    
    // Run the auto migration tool
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
    
    return client
}
