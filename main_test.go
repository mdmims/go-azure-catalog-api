package main

import (
	"github.com/mdmims/go-azure-catalog-api/handler"
	"log"
	"os"
	"testing"
)

var a handler.App

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS asset_type
(
    id SERIAL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    CONSTRAINT asset_type_pk PRIMARY KEY (id)
)`

func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	if _, err := a.DB.Exec("DELETE FROM asset_type"); err != nil {
		log.Fatal(err)
	}

	if _, err := a.DB.Exec("ALTER SEQUENCE IF EXISTS asset_type_seq RESTART WITH 1"); err != nil {
		log.Fatal(err)
	}
}
