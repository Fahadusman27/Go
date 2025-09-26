package middleware

import (
	"strings"

	"tugas/domain/config"
	"tugas/domain/model"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth(userRepo *model.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header required"})
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header format must be Bearer {token}"})
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetJWTSecret()), nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token claims"})
		}

		// ambil user id & role dari claims
		var userID float64
		var role string
		if v, ok := claims["sub"].(float64); ok {
			userID = v
		} else if v, ok := claims["sub"].(int); ok {
			userID = float64(v)
		} else if v, ok := claims["user_id"].(float64); ok {
			userID = v
		} else if v, ok := claims["id"].(float64); ok {
			userID = v
		}

		if r, exists := claims["role"].(string); exists {
			role = r
		}

		if userID != 0 {
			c.Locals("user_id", int(userID))
		}
		c.Locals("role", role)

		return c.Next()
	}
}
