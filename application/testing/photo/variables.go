package photo

import "fmt"

// initial
var (
	dbHost     = "localhost"
	dbName     = "hacktiv"
	dbUser     = "root"
	dbPassword = "root"
	dbPort     = "5432"
	dbTimezone = "Asia/Jakarta"

	dbTestHost     = "localhost"
	dbTestName     = "photos_test"
	dbTestUser     = "root"
	dbTestPassword = "root"
	dbTestPort     = "5432"
	dbTestTimezone = "Asia/Jakarta"

	dbTable = "photos"
)

// connection
var (
	connection     = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", dbHost, dbUser, dbPassword, dbName, dbPort, dbTimezone)
	connectionTest = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", dbTestHost, dbTestUser, dbTestPassword, dbTestName, dbTestPort, dbTestTimezone)
)

// table
var (
	createDatabase  = fmt.Sprintf("CREATE DATABASE %s", dbTestName)
	deleteFromTable = fmt.Sprintf("DELETE FROM %s", dbTable)
	dropTable       = fmt.Sprintf("DROP TABLE %s", dbTable)
	dropaDatabase   = fmt.Sprintf("DROP DATABASE %s", dbTestName)
	createTable     = fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS %s (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        user_id INTEGER NOT NULL,
        description TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP
    )`, dbTable)
)
