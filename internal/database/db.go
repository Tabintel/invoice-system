package database

import (
    "context"
    "log"
    
    "github.com/Tabintel/invoice-system/ent"
    _ "github.com/lib/pq"
)

func NewClient(databaseUrl string) *ent.Client {
    client, err := ent.Open("postgres", databaseUrl)
    if err != nil {
        log.Fatalf("failed opening connection to postgres: %v", err)
    }
    
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
    
    return client
}
