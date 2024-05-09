package surrdb

import (
	"encoding/json"
	"fmt"

	"github.com/surrealdb/surrealdb.go"
)

type Storage struct {
	db     *surrealdb.DB
	Closer Closable
}

type Closable interface {
	Close()
}

func New(host, user, pass, ns, database string) (*Storage, error) {
	db, err := surrealdb.New(host)
	if err != nil {
		return nil, err
	}

	_, err = db.Signin(map[string]interface{}{
		"user": user, "pass": pass,
	})
	if err != nil {
		return nil, err
	}

	if _, err = db.Use(ns, database); err != nil {
		panic(err)
	}

	return &Storage{db, db}, nil
}

func PrintOut[T any](something T) {
	output, _ := json.MarshalIndent(something, "", "  ")

	fmt.Printf("%s \n", output)
}
