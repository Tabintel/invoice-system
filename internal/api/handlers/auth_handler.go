package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/Tabintel/invoice-system/internal/middleware"
    "github.com/Tabintel/invoice-system/internal/service"
    "github.com/Tabintel/invoice-system/internal/ent"
)
type AuthHandler struct {
    userService *service.UserService
    jwtSecret   string
}

func NewAuthHandler(userService *service.UserService, jwtSecret string) *AuthHandler {
    return &AuthHandler{
        userService: userService,
        jwtSecret:   jwtSecret,
    }
}

// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body LoginRequest true "Login credentials"
// @Success 200 {object} LoginResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
    var req LoginRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    user, err := h.userService.Authenticate(c.Context(), req.Email, req.Password)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid credentials",
        })
    }

    token, err := jwt.GenerateToken(user.ID, user.Email, user.Role, h.jwtSecret)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Could not generate token",
        })
    }

    return c.JSON(LoginResponse{
        Token: token,
        User:  user,
    })
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Token string     `json:"token"`
    User  *ent.User `json:"user"`
}
