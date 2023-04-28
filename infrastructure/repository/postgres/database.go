// Package postgres provides the database connection
package postgres

import (
	"gorm.io/gorm"
)

// NewGorm is a function that returns a gorm database connection using  initial configuration
func NewGorm() (gormDB *gorm.DB, err error) {
	var infoDB infoDatabasePostgreSQL
	err = infoDB.getPostgreConn("Databases.PostgreSQL.Localhost")
	if err != nil {
		return nil, err
	}

	gormDB, err = initPostgreDB(gormDB, infoDB)
	if err != nil {
		return nil, err
	}

	var result int
	if err = gormDB.Raw("SELECT 1").Scan(&result).Error; err != nil {
		return nil, err
	}

	return
}
