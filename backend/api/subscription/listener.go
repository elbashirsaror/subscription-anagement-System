package subscription

import (
	"context"
	"ieltsAiTutor/backend/config"
	"log"
	"time"

	"cloud.google.com/go/firestore"
)

func WatchSubscriptions() {
	client := config.FirestoreClient
	if client == nil {
		log.Fatal("❌ FirestoreClient not initialized")
	}

	snapIter := client.Collection("subscriptions").Snapshots(context.Background())
	for {
		snap, err := snapIter.Next()
		if err != nil {
			log.Fatalf("❌ Error reading Firestore updates: %v", err)
		}

		for _, change := range snap.Changes {
			if change.Kind == firestore.DocumentModified {
				userID := change.Doc.Ref.ID
				subscriptionData := change.Doc.Data()
				// ✅ Check expiration when a subscription is modified
				log.Printf("🔄 Subscription expiration checking...")
				log.Printf("🔍 Subscription Data: %+v", subscriptionData)
				if expiresAtRaw, exists := subscriptionData["expires_at"]; exists {
					expiresAt, ok := expiresAtRaw.(time.Time)
					if !ok {
						log.Printf("❌ Unexpected type for expiration date: %T", expiresAtRaw)
						return
					}
					if time.Now().After(expiresAt) {
						log.Printf("❌ Subscription expired for user %s! Revoking access...", userID)
						RevokeUserAccess(userID) // ✅ Call access revocation function
					}
				} else {
					log.Println("⚠️ No expiration date found.")
				}

				// ✅ Check if upgrade request is true and user isn't premium
				log.Printf("🔄 Subscription upgrade request checking...")
				log.Printf("🔍 Subscription Data: %+v", subscriptionData)
				ApproveUpgrade(userID)
				log.Printf("🔄 Subscription updated for user: %s", change.Doc.Ref.ID)
				log.Printf("################################################################")
			}
		}
	}
}
