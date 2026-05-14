package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// DBTX is an interface that allows us to mock the pgxpool.Pool in our unit tests.
// Both *pgxpool.Pool and pgx.Tx implement this interface.
type DBTX interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

// Repository defines the data access methods for the Identity Vault.
type Repository interface {
	UpsertUserCommitment(ctx context.Context, user *User) error
	GetUser(ctx context.Context, discordID string) (*User, error)
}

type pgxRepository struct {
	db DBTX
}

// NewRepository creates a new instance of the PostgreSQL repository.
func NewRepository(db DBTX) Repository {
	return &pgxRepository{db: db}
}

// UpsertUserCommitment securely creates or updates a user's identity commitment in the database.
func (r *pgxRepository) UpsertUserCommitment(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (discord_id, identity_commitment, anchor_status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (discord_id) DO UPDATE
		SET identity_commitment = EXCLUDED.identity_commitment,
			anchor_status = EXCLUDED.anchor_status,
			updated_at = EXCLUDED.updated_at
	`
	
	now := time.Now()
	user.UpdatedAt = now
	if user.CreatedAt.IsZero() {
		user.CreatedAt = now
	}

	_, err := r.db.Exec(ctx, query, user.DiscordID, user.IdentityCommitment, user.AnchorStatus, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("context: failed to upsert user commitment: %w", err)
	}

	return nil
}

// GetUser retrieves a user by their DiscordID from the database.
func (r *pgxRepository) GetUser(ctx context.Context, discordID string) (*User, error) {
	query := `SELECT discord_id, identity_commitment, anchor_status, created_at, updated_at FROM users WHERE discord_id = $1`
	
	var user User
	err := r.db.QueryRow(ctx, query, discordID).Scan(
		&user.DiscordID,
		&user.IdentityCommitment,
		&user.AnchorStatus,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("context: user not found: %w", err)
		}
		return nil, fmt.Errorf("context: failed to get user: %w", err)
	}

	return &user, nil
}
