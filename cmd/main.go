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
    invoiceRepo := repository.NewInvoiceRepository(db)
    userRepo := repository.NewUserRepository(db)

    // Initialize services
    authService := service.NewAuthService(userRepo)
    invoiceService := service.NewInvoiceService(invoiceRepo, userRepo)
    dashboardService := service.NewDashboardService(invoiceRepo, userRepo)
    activityService := service.NewActivityService(invoiceRepo)
    pdfService := service.NewPDFService(invoiceRepo)
    shareService := service.NewShareService(invoiceRepo)

    // Initialize handlers
    authHandler := handlers.NewAuthHandler(authService)
    invoiceHandler := handlers.NewInvoiceHandler(invoiceService)
    dashboardHandler := handlers.NewDashboardHandler(dashboardService)
    activityHandler := handlers.NewActivityHandler(activityService)
    pdfHandler := handlers.NewPDFHandler(pdfService, invoiceService)
    shareHandler := handlers.NewShareHandler(shareService, invoiceService)

    // Router setup
    router := http.NewServeMux()

    // Public routes
    router.HandleFunc("/api/v1/auth/login", authHandler.Login)
    router.HandleFunc("/api/v1/auth/register", authHandler.Register)
    router.HandleFunc("/api/v1/invoices/shared", shareHandler.GetSharedInvoice) // Public shared invoice access

    // Protected routes
    protected := http.NewServeMux()
    protected.HandleFunc("/api/v1/dashboard", dashboardHandler.GetDashboard)
    protected.HandleFunc("/api/v1/invoices", invoiceHandler.CreateInvoice)
    protected.HandleFunc("/api/v1/invoices/pdf", pdfHandler.GenerateInvoicePDF)
    protected.HandleFunc("/api/v1/invoices/share", shareHandler.GenerateShareableLink)
    protected.HandleFunc("/api/v1/activities", activityHandler.GetRecentActivities)

    // Apply JWT middleware to protected routes
    jwtMiddleware := middleware.JWTAuth(os.Getenv("JWT_SECRET"))
    router.Handle("/api/v1/", jwtMiddleware(protected))

    // Apply CORS middleware
    corsHandler := middleware.CORS(router)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Println("Successfully connected to database")
    log.Printf("Server starting on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}