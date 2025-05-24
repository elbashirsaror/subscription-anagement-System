package subscription

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"ieltsAiTutor/backend/config"
)

// ✅ Firestore implementation of SubscriptionRepository
type FirestoreSubscription struct {
	Client *firestore.Client
}

// ✅ Function to add a new subscription
func AddSubscription(userID string, tier string, durationDays int, autoRenew bool) error {
	client := config.FirestoreClient
	if client == nil {
		return fmt.Errorf("❌ Firestore client not initialized")
	}

	// ✅ Calculate expiration date
	expiresAt := time.Now().Add(time.Duration(durationDays) * 24 * time.Hour)

	// ✅ Create subscription object
	newSubscription := Subscription{
		UserID:    userID,
		Tier:      tier,
		ExpiresAt: expiresAt,
		AutoRenew: autoRenew,
		CreatedAt: time.Now(),
	}

	// ✅ Save to Firestore
	docRef := client.Collection("subscriptions").Doc(userID)
	_, err := docRef.Set(context.Background(), newSubscription)
	if err != nil {
		return fmt.Errorf("❌ Failed to create subscription: %v", err)
	}

	log.Printf("✅ Subscription added for user %s (Tier: %s, Expires: %v)", userID, tier, expiresAt)
	return nil
}


func (fs *FirestoreSubscription) GetSubscription(userID string) (*Subscription, error) {
	doc := fs.Client.Collection("subscriptions").Doc(userID)
	snapshot, err := doc.Get(context.Background())
	if err != nil {
		return nil, fmt.Errorf("❌ Failed to fetch subscription: %v", err)
	}

	subscription := &Subscription{}
	if err := snapshot.DataTo(subscription); err != nil {
		return nil, fmt.Errorf("❌ Failed to parse subscription data: %v", err)
	}

	return subscription, nil
}

func (fs *FirestoreSubscription) UpdateExpiration(userID string, newDate time.Time) error {
	doc := fs.Client.Collection("subscriptions").Doc(userID)
	_, err := doc.Update(context.Background(), []firestore.Update{
		{Path: "expires_at", Value: newDate},
	})
	if err != nil {
		return fmt.Errorf("❌ Failed to update expiration: %v", err)
	}
	log.Printf("✅ Subscription expiration updated for user %s", userID)
	return nil
}