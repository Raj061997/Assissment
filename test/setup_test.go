package test

import (
	"log"
	"os"
	"testing"

	"example/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Test Database Connection
func TestMain(m *testing.M) {
	dsn := "host=localhost user=postgres password=ADMIN dbname=postgres port=5432 sslmode=disable"
	var err error
	database.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	// Run tests
	exitCode := m.Run()
	os.Exit(exitCode)
}
