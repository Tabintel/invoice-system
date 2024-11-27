package main

import (
    "log"
    "net/http"
    "os"

    "github.com/joho/godotenv"
    "github.com/Tabintel/invoice-system/internal/config"
    "github.com/Tabintel/invoice-system/internal/api/handlers"
    "github.com/Tabintel/invoice-system/internal/service"
    "github.com/Tabintel/invoice-system/internal/repository"
    "github.com/Tabintel/invoice-system/internal/middleware"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Printf("Warning: .env file not found")
    }

    db, err := config.NewDatabase(os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatal("cannot connect to db:", err)
    }
    defer db.Close()

    // Initialize repositories
    repo := repository.NewRepository(db)

    // Initialize services
    authService := service.NewAuthService(repo)
    invoiceService := service.NewInvoiceService(repo, repo, nil)
    dashboardService := service.NewDashboardService(repo)
    activityService := service.NewActivityService(repo)
    pdfService := service.NewPDFService(repo)
    notificationService := service.NewNotificationService(
        os.Getenv("SMTP_EMAIL"),
        os.Getenv("SMTP_PASSWORD"),
        os.Getenv("SMTP_HOST"),
        os.Getenv("SMTP_PORT"),
    )

    // Initialize handlers
    authHandler := handlers.NewAuthHandler(authService)
    invoiceHandler := handlers.NewInvoiceHandler(invoiceService)
    dashboardHandler := handlers.NewDashboardHandler(dashboardService)
    activityHandler := handlers.NewActivityHandler(activityService)
    pdfHandler := handlers.NewPDFHandler(pdfService, invoiceService)

    // Middleware
    jwtMiddleware := middleware.JWTAuth(os.Getenv("JWT_SECRET"))
    
    // Router setup
    router := http.NewServeMux()

    // Public routes
    router.HandleFunc("/api/v1/auth/login", authHandler.Login)
    router.HandleFunc("/api/v1/auth/register", authHandler.Register)
    router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })

    // Protected routes
    protected := http.NewServeMux()
    protected.HandleFunc("/api/v1/dashboard", dashboardHandler.GetDashboard)
    protected.HandleFunc("/api/v1/invoices", invoiceHandler.CreateInvoice)
    protected.HandleFunc("/api/v1/invoices/pdf", pdfHandler.GenerateInvoicePDF)
    protected.HandleFunc("/api/v1/activities", activityHandler.GetRecentActivities)

    // Apply JWT middleware to protected routes
    router.Handle("/api/v1/", jwtMiddleware(protected))

    // CORS middleware
    corsHandler := middleware.CORS(router)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    server := &http.Server{
        Addr:    ":" + port,
        Handler: corsHandler,
    }

    //log.Println("Successfully connected to database")
    log.Printf("Server starting on port %s", port)
    log.Fatal(server.ListenAndServe())
}