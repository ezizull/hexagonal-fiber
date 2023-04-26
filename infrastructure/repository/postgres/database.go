// Package postgres provides the database connection
package postgres

import (
	"gorm.io/gorm"
)

// NewGorm is a function that returns a gorm database connection using  initial configuration
func NewGorm() (gormDB *gorm.DB, err error) {
	var infoPg infoDatabasePostgreSQL
	err = infoPg.getPostgreConn("Databases.PostgreSQL.Localhost")
	if err != nil {
		return nil, err
	}

	gormDB, err = initPostgreDB(gormDB, infoPg)
	if err != nil {
		return nil, err
	}

	var result int
	// Test the connection by executing a simple query
	if err = gormDB.Raw("SELECT 1").Scan(&result).Error; err != nil {
		return nil, err
	}

	return
}
