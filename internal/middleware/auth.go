package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func EnsureAuthenticated() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authorization := ctx.Get("Authorization")

		if len(authorization) < 8 {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Missing authentication token"})
		}

		tokenString := authorization[7:]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("JWT_SECRET"), nil
		})

		if err != nil {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token type"})
		}

		if !token.Valid {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Expired token. Please login again"})
		}

		return ctx.Next()
	}
}
