package Database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)


var db *sql.DB

type Database struct {
	connectionString string
	driver string
}

func Get() *sql.DB {
	return connection(&Database{connectionString: databaseFile(), driver: "libsql"})
}

func connection(database *Database) *sql.DB {
	if db != nil {
		return db
	}

	var err error
	db, err = sql.Open(database.driver, database.connectionString)
	if err!= nil {
		panic(err)
	}

	db.SetMaxOpenConns(10)

	return db
}

func databaseFile() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dirname := fmt.Sprintf("%s/.notihub", homeDir)
	fileLocation := fmt.Sprintf("%s/db.sqlite3", dirname)

	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		os.Mkdir(dirname, 0755)
	}

	if _, err := os.Stat(fileLocation); os.IsNotExist(err) {
		file, err := os.Create(fileLocation)
		if err != nil {
			panic(err)
		}
		file.Close()
	}

	return fmt.Sprintf("file://%s", fileLocation)
}
