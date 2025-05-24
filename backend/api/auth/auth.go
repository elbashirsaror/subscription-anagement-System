package auth

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/elbashirsaror/subscription-management-system/backend/config"
	"firebase.google.com/go/auth"
)

type UserSignup struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(c *fiber.Ctx) error {
	if config.FirebaseAuth == nil {
		log.Println("❌ FirebaseAuth is nil!")
		return c.Status(500).JSON(fiber.Map{"error": "FirebaseAuth not initialized"})
	}
	var data UserSignup
	if err := c.BodyParser(&data); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input format"})
	}

	// ✅ Ensure FirebaseAuth is initialized
	if config.FirebaseAuth == nil {
		log.Println("FirebaseAuth is nil!")
		return c.Status(500).JSON(fiber.Map{"error": "FirebaseAuth not initialized"})
	}

	params := (&auth.UserToCreate{}).Email(data.Email).Password(data.Password)
	user, err := config.FirebaseAuth.CreateUser(c.Context(), params)
	if err != nil {
		log.Printf("Firebase CreateUser failed: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "User registered", "user_id": user.UID})
}

func LoginUser(c *fiber.Ctx) error {
	if config.FirebaseAuth == nil {
		return c.Status(500).JSON(fiber.Map{"error": "FirebaseAuth not initialized"})
	}

	var data UserSignup
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	token, err := config.FirebaseAuth.CustomToken(c.Context(), data.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"token": token})
}