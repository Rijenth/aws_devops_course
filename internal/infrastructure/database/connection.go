package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rijenth/gRPC/internal/config"
)

func InitDB() (*sql.DB, error) {
	cfg, err := config.LoadDatabaseConfig()

	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, fmt.Errorf("failed to open the database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	log.Println("Database connection established")

	if err := runMigrations(db); err != nil {
		return nil, fmt.Errorf("failed to apply migrations: %v", err)
	}

	return db, nil
}

func runMigrations(db *sql.DB) error {
	migrationPath := "./internal/infrastructure/database/migrations"

	files, err := os.ReadDir(migrationPath)
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %v", err)
	}

	var migrationFiles []string

	for _, migration := range files {
		if !migration.IsDir() {
			migrationFiles = append(migrationFiles, migration.Name())
		}
	}

	sort.Strings(migrationFiles)

	for _, file := range migrationFiles {
		filePath := filepath.Join(migrationPath, file)
		log.Printf("Applying migration: %s\n", filePath)

		sqlQuery, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %v", file, err)
		}

		_, err = db.Exec(string(sqlQuery))
		if err != nil {
			return fmt.Errorf("failed to execute migration %s: %v", file, err)
		}
	}

	log.Println("All migrations applied successfully")
	return nil
}
