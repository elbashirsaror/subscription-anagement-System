package subscription


func ValidateSubscription(userID string) (bool, error) {
	client := config.FirestoreClient
	if client == nil {
		return false, fmt.Errorf("Firestore client not initialized")
	}

	// ✅ Fetch subscription data
	doc := client.Collection("subscriptions").Doc(userID)
	snapshot, err := doc.Get(context.Background())
	if err != nil {
		return false, fmt.Errorf("Failed to fetch subscription: %v", err)
	}

	// ✅ Check subscription status
	var subscriptionData map[string]interface{}
	if err := snapshot.DataTo(&subscriptionData); err != nil {
		return false, fmt.Errorf("Error processing subscription data: %v", err)
	}

	// ✅ Validate user subscription level
	tier, exists := subscriptionData["tier"]
	if !exists || tier == "free" {
		return false, nil
	}
	
	// ✅ Check expiration
	expiresAtStr, expiresExists := subscriptionData["expires_at"].(string) // ✅ Properly cast to string
	if !exists || !expiresExists || tier == "free" {
		return false, nil
	}
	expiresAt, err := time.Parse(time.RFC3339, expiresAtStr)
	if err != nil {
		return false, fmt.Errorf("Invalid expiration format")
	}

	if time.Now().After(expiresAt) {
		log.Println("❌ Subscription has expired!")
		return false, nil
	}



	log.Println("✅ Subscription is valid!")
	return true, nil
}