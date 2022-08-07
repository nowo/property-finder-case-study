package testdata

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	// CREATE - operation to create database
	CREATE = "CREATE"
	// DROP - operation to drop database
	DROP = "DROP"
)

// MockedDB is used in unit tests to mock db
func MockedDB(operation string) {
	/*
	   If tests are running in CI, environment variables should not be loaded.
	   The reason is environment vars will be provided through CI config file.
	*/

	dbName := os.Getenv("POSTGRES_DB")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASS")

	// createdb => https://www.postgresql.org/docs/7.0/app-createdb.htm
	// dropdb => https://www.postgresql.org/docs/7.0/app-dropdb.htm
	var command string

	if operation == CREATE {
		command = "createdb"
	} else {
		command = "dropdb"
	}

	// createdb & dropdb commands have same configuration syntax.
	cmd := exec.Command(command, "-h", "localhost", "-U", pgUser, "-e", dbName)
	cmd.Env = os.Environ()

	/*
	   if we normally execute createdb/dropdb, we will be propmted to provide password.
	   To inject password automatically, we have to set PGPASSWORD as prefix.
	*/
	cmd.Env = append(cmd.Env, fmt.Sprintf("PGPASSWORD=%v", pgPassword))

	if err := cmd.Run(); err != nil {
		log.Fatalf("Error executing %v on %v.\n%v", command, dbName, err)
	}

	/*
	   Alternatively instead of createdb/dropdb, you can use
	   psql -c "CREATE/DROP DATABASE DBNAME" "DATABASE_URL"
	*/
}
