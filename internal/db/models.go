package db

import "time"

// User represents a verified identity commitment within the Identity Vault.
// Notice that no PII (Personally Identifiable Information) such as names,
// emails, or clear-text ID numbers are stored in this model.
type User struct {
	// DiscordID is the unique identifier for the user on Discord.
	DiscordID string
	// IdentityCommitment is the Poseidon hash of the UserSecret and AnchorID.
	IdentityCommitment string
	// AnchorStatus represents the verification tier (e.g., 'VERIFIED', 'PENDING').
	AnchorStatus string
	// CreatedAt is the timestamp when the user was first registered.
	CreatedAt time.Time
	// UpdatedAt is the timestamp of the last update to the user's record.
	UpdatedAt time.Time
}
