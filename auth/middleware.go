package auth

import (
	"absensi_karyawan/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *fiber.Ctx) error {
	auth := c.Get("Authorization")

	if auth == "" {
		return c.Status(401).JSON(fiber.Map{"error": "missing token"})
	}

	// split "Bearer token"
	parts := strings.Split(auth, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(401).JSON(fiber.Map{"error": "invalid token format"})
	}

	tokenStr := parts[1]

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return utils.SECRET, nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "invalid token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.SendStatus(401)
	}

	// simpan ke context
	c.Locals("user_id", claims["user_id"])
	c.Locals("role", claims["role"])

	return c.Next()
}

func AdminOnly(c *fiber.Ctx) error {
	role, ok := c.Locals("role").(string)
	if !ok || role != "admin" {
		return c.Status(403).JSON(fiber.Map{"error": "admin only"})
	}
	return c.Next()
}
