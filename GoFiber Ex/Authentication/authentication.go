package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("secret")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	app := fiber.New()

	app.Post("/login", login)
	app.Post("/logout", logout)
	app.Get("/protected", jwtMiddleware(), protected)

	log.Fatal(app.Listen(":3036"))
}

func login(c *fiber.Ctx) error {
	var creds Credentials
	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	// Dummy user validation
	if creds.Username != "user" || creds.Password != "pass" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": creds.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not generate token"})
	}

	return c.JSON(fiber.Map{"token": tokenString})
}

func logout(c *fiber.Ctx) error {
	// Invalidate the token (optional: implement token blacklist)
	return c.SendStatus(fiber.StatusOK)
}

func jwtMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing or invalid token"})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "invalid signing method")
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}

		c.Locals("user", token.Claims.(jwt.MapClaims)["username"])
		return c.Next()
	}
}

func protected(c *fiber.Ctx) error {
	username := c.Locals("user").(string)
	return c.JSON(fiber.Map{"message": "Welcome " + username})
}
