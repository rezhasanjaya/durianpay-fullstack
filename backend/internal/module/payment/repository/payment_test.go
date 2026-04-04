package repository

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestPaymentRepo_GetPayments(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE payments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			merchant TEXT NOT NULL,
			created_at TEXT NOT NULL,
			amount TEXT NOT NULL,
			status TEXT NOT NULL
		);
	`)
	if err != nil {
		t.Fatalf("create table: %v", err)
	}

	_, err = db.Exec(`
		INSERT INTO payments (merchant, status, amount, created_at)
		VALUES
			('Merchant A', 'completed', '10000', '2026-04-01T10:00:00Z'),
			('Merchant B', 'failed', '20000', '2026-04-01T11:00:00Z');
	`)
	if err != nil {
		t.Fatalf("seed: %v", err)
	}

	repo := NewPaymentRepo(db)
	res, err := repo.GetPayments("", "")
	if err != nil {
		t.Fatalf("GetPayments error: %v", err)
	}

	if len(res) != 2 {
		t.Fatalf("expected 2, got %d", len(res))
	}
}
