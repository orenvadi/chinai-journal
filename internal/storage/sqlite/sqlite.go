package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// we need that for dependency inversion
type Storage struct {
	db *sql.DB
}

const (
	pathToCreateTablesSQL = "/home/orenvady/Repos/go/chinai-journal/internal/storage/sqlite/create_tables.sql"
)

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// https://ya.ru/woreijgeroijfsoerjgfeoirfjoeirjfoerij
	// https://my.sh/sqlite

	queries, err := QueryFromFile(pathToCreateTablesSQL)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for i, query := range queries {
		fmt.Println("Created table ", i)
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Database tables created successfully.")

	return &Storage{db: db}, nil
}

func QueryFromFile(path string) ([]string, error) {
	// returns list of queries to create tables
	fileContent, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return []string{}, err
	}

	contentString := string(fileContent)

	queries := strings.Split(contentString, "-- ")

	return queries, nil
}
