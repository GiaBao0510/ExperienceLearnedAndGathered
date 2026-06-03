package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Init khởi tạo kết nối đến cơ sở dữ liệu
func Init(dbPath string) *sql.DB {
	database, err := sql.Open("postgres", dbPath)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	//  Kiểm tra kết nối
	if err := database.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	migrate(database)
	log.Println("Kết nối đến database thành công và đã thực hiện migration.")
	return database
}

// Tạo Migration để tạo bảng
func migrate(db *sql.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS api_keys (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		user_id INTEGER NOT NULL,
		name VARCHAR(255) NOT NULL,
		key_hint VARCHAR(10) NOT NULL,
		key_hash VARCHAR(255) NOT NULL UNIQUE,
		scopes TEXT[] DEFAULT '{}',
		rate_limit INTEGER NOT NULL DEFAULT 1000,
		is_active BOOLEAN NOT NULL DEFAULT TRUE,
		last_used TIMESTAMPTZ,
		expires_at TIMESTAMPTZ,
		created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_api_key_hash ON api_keys(key_hash);
	`

	if _, err := db.Exec(schema); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
