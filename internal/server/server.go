package server

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"
    "github.com/Tabintel/invoice-system/ent"
    "github.com/Tabintel/invoice-system/internal/services"
    "github.com/Tabintel/invoice-system/internal/handlers"
)
type Server struct {
    router *chi.Mux
    db     *ent.Client
}

func NewServer(db *ent.Client) *Server {
    s := &Server{
        router: chi.NewRouter(),
        db:     db,
    }
    
    s.setupRoutes()
    return s
}

func (s *Server) setupRoutes() {
    s.router.Use(middleware.Logger)
    s.router.Use(middleware.Recoverer)
    s.router.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:          300,
    }))

    s.router.Get("/health", s.handleHealth())
    invoiceService := services.NewInvoiceService(s.db)
    invoiceHandler := handlers.NewInvoiceHandler(invoiceService)
    
    s.router.Post("/api/invoices", invoiceHandler.Create())

    s.router.Get("/api/invoices", invoiceHandler.List())

    s.router.Put("/api/invoices/{id}/status", invoiceHandler.UpdateStatus())

}

func (s *Server) handleHealth() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    }
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    s.router.ServeHTTP(w, r)
}

    
