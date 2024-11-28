package server

import (
    //"encoding/json"
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"
    "github.com/Tabintel/invoice-system/ent"
    //"github.com/Tabintel/invoice-system/internal/docs"
    "github.com/Tabintel/invoice-system/internal/services"
    "github.com/Tabintel/invoice-system/internal/handlers"
   // httpSwagger "github.com/swaggo/http-swagger"
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

        // Serve Swagger UI
        s.router.Get("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
            http.ServeFile(w, r, "internal/api/swagger.json")
        })

        //s.router.Handle("/docs/*", http.StripPrefix("/docs/", http.FileServer(http.Dir("swagger-ui"))))

         // Swagger UI setup
         //s.router.Get("/swagger/*", httpSwagger.Handler(
           // httpSwagger.URL("/swagger/doc.json"),
        //))

        customerService := services.NewCustomerService(s.db)
        customerHandler := handlers.NewCustomerHandler(customerService)
    
        s.router.Post("/api/customers", customerHandler.Create())
        s.router.Get("/api/customers", customerHandler.List())
        s.router.Get("/api/customers/{id}", customerHandler.Get())
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

    

func (s *Server) Router() *chi.Mux {
    return s.router
}
