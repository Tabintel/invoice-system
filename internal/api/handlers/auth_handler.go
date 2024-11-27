package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/Tabintel/invoice-system/internal/service"
)

type AuthHandler struct {
    service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
    return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req service.LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    token, err := h.service.Login(r.Context(), &req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req service.RegisterRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Validate required fields for invoice system
    if req.Name == "" || req.Email == "" || req.Password == "" || req.CompanyName == "" {
        http.Error(w, "Missing required fields", http.StatusBadRequest)
        return
    }

    if err := h.service.Register(r.Context(), &req); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Auto-login after registration
    token, err := h.service.Login(r.Context(), &service.LoginRequest{
        Email:    req.Email,
        Password: req.Password,
    })
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "token": token,
        "message": "Registration successful",
    })
}
