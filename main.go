package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"ieltsAiTutor/backend/config"
	"ieltsAiTutor/backend/api/auth"
	"ieltsAiTutor/backend/api/subscription/listener" // ✅ Add this import

)

func main() {
	app := fiber.New()

	// ✅ Ensure Firebase initializes successfully
	err := config.InitFirebase()
	if err != nil {
		log.Fatalf("Firebase failed to initialize: %v", err)
	}

	err = config.InitFirestore()
	if err != nil {
		log.Fatal("❌ Firestore failed to initialize!")
	}
	
	// ✅ Start Firestore subscription monitoring (continuous in production)
	go listener.WatchSubscriptions()

	// Auth Routes
	app.Post("/signup", auth.RegisterUser)
	app.Post("/login", auth.LoginUser)

	log.Println("Server running on port 3000...")
	app.Listen(":3000")
}