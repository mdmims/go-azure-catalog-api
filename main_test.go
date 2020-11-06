package main

import (
	"github.com/mdmims/go-azure-catalog-api/handler"
	"log"
	"net/http"
	"net/http/httptest"
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

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestMain(m *testing.M) {
	// load env vars
	initEnv()

	a.Initialize(
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"),
		os.Getenv("TEST_DB_PORT"))

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/products", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}
