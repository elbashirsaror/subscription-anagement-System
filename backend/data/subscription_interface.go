package subscription

import "time"

// ✅ SubscriptionRepository defines subscription-related operations
type SubscriptionRepository interface {
	AddSubscription(userID string, tier string, durationDays int, autoRenew bool) error
	GetSubscription(userID string) (*Subscription, error)
	UpdateExpiration(userID string, newDate time.Time) error
	RequestUpgrade(userID string) error
	ApproveUpgrade(userID string) error
	RenewSubscription(userID string) error
}

