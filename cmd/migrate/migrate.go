package migrate

import "gorm.io/gorm"

var PostgresDB *gorm.DB

// SetMigrateDB sets the postgresDB instance for the cmd package
func SetMigrateDB(postgres *gorm.DB) {
	PostgresDB = postgres
}
