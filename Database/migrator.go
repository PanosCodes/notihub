package Database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"slices"
)

func Migrate() {
	db := Get()
	createMigrationTable(db)

	entries, err := os.ReadDir(migrationsPath())

	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		runMigrationIfNecessary(db, entry.Name())
	}
}

func runMigrationIfNecessary(db *sql.DB, name string) {
	existingMigrations := getExistingMigrations(db)

	if slices.Contains(existingMigrations, name) {
		return
	}

	runMigration(db, name)
	insertMigration(db, name)
}

func runMigration(db *sql.DB, name string) {
	fmt.Println("Running migration", name)
	migrationPath := fmt.Sprintf("%s/%s", migrationsPath(), name)

	f, err := os.ReadFile(migrationPath)
	if err != nil {
		panic(err)
	}

	sql := string(f)

	_, err = db.Exec(sql)
	if err!= nil {
		panic(err)
	}
}



func createMigrationTable(db *sql.DB ) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS migrations (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL)")
	if err!= nil {
		panic(err)
	}
}

func insertMigration(db *sql.DB, name string) {
	_, err := db.Exec("INSERT INTO migrations (name) VALUES (?)", name)
	if err!= nil {
		panic(err)
	}
}

func getExistingMigrations(db *sql.DB) []string {
	var existingMigrations []string

	rows, err := db.Query("SELECT name FROM migrations")
	if err!= nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var name string
		rows.Scan(&name)

		existingMigrations = append(existingMigrations, name)
	}

	return existingMigrations
}

func basePath() string {
	_, b, _, _ := runtime.Caller(0)

    return filepath.Dir(b)
}

func migrationsPath() string {
	return fmt.Sprintf("%s/migrations", basePath())
}
