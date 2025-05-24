package subscription

import (
	"context"
	"fmt"
	"github.com/elbashirsaror/subscription-management-system/backend/config"
	"log"

	"cloud.google.com/go/firestore"
)

func RequestUpgrade(userID string) error {
	client := config.FirestoreClient
	if client == nil {

		return fmt.Errorf("Firestore client not initialized")
	}

	doc := client.Collection("subscriptions").Doc(userID)
	// upgrade to requested to true
	_, err := doc.Update(context.Background(), []firestore.Update{
		{Path: "upgrade_requested", Value: true},
	})

	if err != nil {
		return fmt.Errorf("Failed to request upgrade for user %s: %v", userID, err)
	}

	log.Printf("✅ Upgrade request recorded for user %s", userID)
	return nil
}

// Allow admins to Approve Upgrades
func ApproveUpgrade(userID string) error {
	client := config.FirestoreClient
	if client == nil {
		return fmt.Errorf("Firestore client not initialized")
	}

	doc := client.Collection("subscriptions").Doc(userID)
	snapshot, err := doc.Get(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to fetch user data: %v", err)
	}

	upgradeRequested, exists := snapshot.Data()["upgrade_requested"].(bool)
	// ✅ Validate user subscription level
	tier, exists := snapshot.Data()["tier"].(string)

	if !exists || !upgradeRequested {
		log.Printf("❌ Upgrade request not found for user %s", userID)
		return fmt.Errorf("Upgrade request not found for user %s", userID)
	}
	if !exists || tier == "premium" {
		log.Printf("❌ User has an active premium subscription %s", userID)
		_, err = doc.Update(context.Background(), []firestore.Update{
			{Path: "upgrade_requested", Value: false}, // ✅ Reset flag
		})
		return fmt.Errorf("User has an active premium subscription %s", userID)
	}

	_, err = doc.Update(context.Background(), []firestore.Update{
		{Path: "tier", Value: "premium"},
		{Path: "upgrade_requested", Value: false}, // ✅ Reset flag
	})

	if err != nil {
		return fmt.Errorf("Failed to upgrade user %s to premium: %v", userID, err)
	}

	log.Printf("✅ User %s successfully upgraded to premium!", userID)
	return nil
}
