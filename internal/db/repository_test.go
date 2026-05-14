package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// mockDBTX is a simple mock implementation of the DBTX interface for testing.
type mockDBTX struct {
	execErr  error
	queryErr error
	row      pgx.Row
}

func (m *mockDBTX) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	return pgconn.CommandTag{}, m.execErr
}

func (m *mockDBTX) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return nil, m.queryErr
}

func (m *mockDBTX) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return m.row
}

// mockRow is a mock implementation of pgx.Row.
type mockRow struct {
	scanErr error
}

func (m *mockRow) Scan(dest ...any) error {
	return m.scanErr
}

func TestUpsertUserCommitment(t *testing.T) {
	mockDB := &mockDBTX{
		execErr: nil, // Simulating a successful exec
	}
	repo := NewRepository(mockDB)

	user := &User{
		DiscordID:          "123456789012345678",
		IdentityCommitment: "0xPoseidonHashExample123",
		AnchorStatus:       "VERIFIED",
	}

	err := repo.UpsertUserCommitment(context.Background(), user)
	if err != nil {
		t.Errorf("expected no error on valid upsert, got %v", err)
	}

	if user.CreatedAt.IsZero() {
		t.Errorf("expected CreatedAt to be automatically set by UpsertUserCommitment, but it was zero")
	}
	if user.UpdatedAt.IsZero() {
		t.Errorf("expected UpdatedAt to be automatically set by UpsertUserCommitment, but it was zero")
	}
}

func TestGetUser_Success(t *testing.T) {
	mockDB := &mockDBTX{
		row: &mockRow{
			scanErr: nil, // Simulating successful row scan
		},
	}
	repo := NewRepository(mockDB)

	user, err := repo.GetUser(context.Background(), "123456789012345678")
	if err != nil {
		t.Errorf("expected no error on valid get, got %v", err)
	}
	if user == nil {
		t.Errorf("expected user to be returned, got nil")
	}
}

func TestGetUser_NotFound(t *testing.T) {
	mockDB := &mockDBTX{
		row: &mockRow{
			scanErr: pgx.ErrNoRows, // Simulating user not found
		},
	}
	repo := NewRepository(mockDB)

	user, err := repo.GetUser(context.Background(), "non_existent_id")
	if err == nil {
		t.Errorf("expected error when user not found, got nil")
	}
	if user != nil {
		t.Errorf("expected nil user when not found, got %v", user)
	}
}
