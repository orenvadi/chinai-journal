package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/surrealdb/surrealdb.go"
)

type migrationResult struct {
	Result string      `json:"result"`
	Status interface{} `json:"status,omitempty"`
	Time   string      `json:"time"`
}

// for Args
var (
	storageUser        string
	storagePassword    string
	storageHost        string
	storageDbName      string
	storageDbNameSpace string
	migrationsPath     string
)

func main() {
	// parsing cli args

	parseArgs()

	// opening and closing surrealdb connection
	db, err := surrealdb.New(storageHost)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// sign in to db
	if _, err = db.Signin(map[string]interface{}{
		"user": storageUser, "pass": storagePassword,
	}); err != nil {
		panic(err)
	}

	if _, err = db.Use(storageDbNameSpace, storageDbName); err != nil {
		panic(err)
	}

	// get schema from file
	schema, err := readFileAll(migrationsPath)
	if err != nil {
		panic(err)
	}

	// run schema

	rawResult, err := db.Query(schema, map[string]interface{}{})
	if err != nil {
		panic(err)
	}

	// unmarshall result

	result := make([]migrationResult, 60)

	if err = surrealdb.Unmarshal(rawResult, &result); err != nil {
		panic(err)
	}

	// print result

	for i, res := range result {
		fmt.Println(i, res.Status)
	}
}

func parseArgs() {
	flag.StringVar(&storageUser, "user", "root", "Surrealdb user")
	flag.StringVar(&storagePassword, "pass", "root", "Surrealdb user")
	flag.StringVar(&storageHost, "host", "", "Surrealdb host")
	flag.StringVar(&storageDbName, "dbname", "", "Surrealdb name")
	flag.StringVar(&storageDbNameSpace, "dbnamespace", "", "Surrealdb namespace")

	flag.StringVar(&migrationsPath, "migr-path", "", "path to migrations")
	flag.Parse()

	switch {
	case storageUser == "":
		panic("user is required")
	case storagePassword == "":
		panic("pass is required")
	case storageHost == "":
		panic("host is required")
	case storageDbName == "":
		panic("name is required")
	case storageDbNameSpace == "":
		panic("namespace is required")

	case migrationsPath == "":
		panic("migr-path is required")
	}
}

func readFileAll(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
