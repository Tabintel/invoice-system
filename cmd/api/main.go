package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/Tabintel/invoice-system/internal/config"
    "github.com/Tabintel/invoice-system/internal/api/handlers"
    "github.com/Tabintel/invoice-system/internal/repository"
    "github.com/Tabintel/invoice-system/internal/service"
    "github.com/Tabintel/invoice-system/internal/middleware"
)

func main() {
    // Load config
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal("Cannot load config:", err)
    }

    // Initialize database
    db := repository.NewDatabase(cfg.GetDBConnString())
    defer db.Close()

    // Initialize repositories
    invoiceRepo := repository.NewInvoiceRepository(db)
    userRepo := repository.NewUserRepository(db)
    activityRepo := repository.NewActivityRepository(db)
    dashboardRepo := repository.NewDashboardRepository(db)

    // Initialize services
    invoiceService := service.NewInvoiceService(invoiceRepo, activityRepo)
    userService := service.NewUserService(userRepo)
    dashboardService := service.NewDashboardService(dashboardRepo)
    activityService := service.NewActivityService(activityRepo)

    // Initialize handlers
    authHandler := handlers.NewAuthHandler(userService, cfg.JWTSecret)
    invoiceHandler := handlers.NewInvoiceHandler(invoiceService)
    dashboardHandler := handlers.NewDashboardHandler(dashboardService)
    activityHandler := handlers.NewActivityHandler(activityService)

    // Setup Fiber app
    app := fiber.New()
    app.Use(cors.New())
    app.Use(logger.New())

    // API routes
    api := app.Group("/api/v1")

    // Public routes
    api.Post("/auth/login", authHandler.Login)
    api.Post("/auth/register", authHandler.Register)

    // Protected routes
    protected := api.Use(middleware.JWTProtected(cfg.JWTSecret))
    
    // Dashboard routes
    protected.Get("/dashboard", dashboardHandler.GetStats)
    
    // Invoice routes
    protected.Post("/invoices", invoiceHandler.CreateInvoice)
    protected.Get("/invoices", invoiceHandler.ListInvoices)
    protected.Get("/invoices/:id", invoiceHandler.GetInvoice)
    protected.Put("/invoices/:id/status", invoiceHandler.UpdateStatus)
    protected.Get("/invoices/:id/pdf", invoiceHandler.GeneratePDF)
    protected.Get("/invoices/:id/share", invoiceHandler.GenerateShareableLink)
    
    // Activity routes
    protected.Get("/activities", activityHandler.GetRecentActivities)

    log.Fatal(app.Listen(":"+cfg.ServerPort))
}
