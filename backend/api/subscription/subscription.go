package subscription

import (
	"ieltsAiTutor/backend/api/subscription/revoke" // ✅ Import separated functions
	"log"
	"time"
)

func ProcessSubscription(userID string) {
	// ✅ Check expiration
	if time.Now().After(GetExpirationDate(userID)) {
		log.Printf("❌ Subscription expired for user %s! Revoking access...", userID)
		revoke.RevokeUserAccess(userID) // ✅ Calls function from revoke.go
	}
}

// real-time listeners

/**
Add Subscription Upgrade logic

*/
