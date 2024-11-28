package handlers

import (
    "encoding/json"
    "net/http"
    "log"
    "github.com/Tabintel/invoice-system/internal/services"
)

type CustomerHandler struct {
    service *services.CustomerService
}

func NewCustomerHandler(service *services.CustomerService) *CustomerHandler {
    return &CustomerHandler{service: service}
}

func (h *CustomerHandler) Create() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var input services.CreateCustomerInput
        if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
            log.Printf("Error decoding request: %v", err)
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }

        customer, err := h.service.CreateCustomer(r.Context(), input)
        if err != nil {
            log.Printf("Error creating customer: %v", err)
            http.Error(w, "Failed to create customer", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "status": "success",
            "data":   customer,
        })
    }
}

func (h *CustomerHandler) List() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        customers, err := h.service.ListCustomers(r.Context())
        if err != nil {
            log.Printf("Error listing customers: %v", err)
            http.Error(w, "Failed to list customers", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "status": "success",
            "data":   customers,
        })
    }
}
