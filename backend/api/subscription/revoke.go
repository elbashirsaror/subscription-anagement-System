package subscription

import (
	"context"
	"github.com/elbashirsaror/subscription-management-system/backend/config"
	"log"
	"time"

	"cloud.google.com/go/firestore"
)

func RevokeUserAccess(userID string) {
	client := config.FirestoreClient
	if client == nil {
		log.Fatal("❌ FirestoreClient not initialized")
	}

	// ✅ Update user’s subscription tier to "free"
	doc := client.Collection("subscriptions").Doc(userID)

	// Get the current time
	currentTime := time.Now()

	// Increase time by 50 years
	futureTime := currentTime.AddDate(50, 0, 0) // Adding 50 years

	// Store in a timestamp variable (Unix timestamp)

	_, err := doc.Update(context.Background(), []firestore.Update{
		{Path: "tier", Value: "free"},
	})
	// Update Firestore document
	_, err = doc.Update(context.Background(), []firestore.Update{
		{Path: "expires_at", Value: futureTime},
	})

	if err != nil && err != nil {
		log.Printf("❌ Failed to revoke access for user %s: %v", userID, err)
	} else {
		//log.Printf("✅ Access revoked for user %s. Now on 'free' tier.", userID)
		log.Printf("✅ timestamp updated for user %s. ", userID)
	}
}
