package middleware

import (
    "github.com/gofiber/fiber/v2"
    jwtware "github.com/gofiber/jwt/v3"
    "github.com/golang-jwt/jwt/v4"
)

func JWTProtected(secret string) fiber.Handler {
    return jwtware.New(jwtware.Config{
        SigningKey:    []byte(secret),
        ErrorHandler:  jwtError,
    })
}

func jwtError(c *fiber.Ctx, err error) error {
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
        "error": "Unauthorized",
        "message": "Invalid or expired token",
    })
}
