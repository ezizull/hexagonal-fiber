package photo

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func setupDatabase(its *IntTestSuite, db *gorm.DB) *gorm.DB {
	its.T().Log("setting up database")

	tx := db.Exec(createDatabase)
	if tx.Error != nil {
		its.FailNowf("unable to create database", tx.Error.Error())
	}

	var err error
	db, err = gorm.Open(postgres.Open(connectionTest), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		its.FailNowf("unable to connect to database", err.Error())
	}

	tx = db.Exec(createTable)
	if tx.Error != nil {
		its.FailNowf("unable to create table", tx.Error.Error())
	}

	return db
}

func seedTestTable(its *IntTestSuite, db *gorm.DB) {
	its.T().Log("seeding test table")

	for i := 1; i <= 2; i++ {
		query := fmt.Sprintf("INSERT INTO books (id, title, user_id, description, created_at, updated_at, deleted_at) VALUES (%d, 'Book %d', %d, 'Description %d', NOW(), NOW(), NULL)", i, i, i, i)
		tx := db.Exec(query)
		if tx.Error != nil {
			its.FailNowf("unable to seed table", tx.Error.Error())
		}
	}
}

func cleanTable(its *IntTestSuite) {
	its.T().Log("cleaning database")

	tx := its.db.Exec(deleteFromTable)
	if tx.Error != nil {
		its.FailNowf("unable to clean table", tx.Error.Error())
	}
}

func tearDownDatabase(its *IntTestSuite) {
	its.T().Log("tearing down database")

	tx := its.db.Exec(dropTable)
	if tx.Error != nil {
		its.FailNowf("unable to drop table", tx.Error.Error())
	}

	db, err := its.db.DB()
	if err != nil {
		its.FailNowf("unable to close database", err.Error())
	}

	err = db.Close()
	if err != nil {
		its.FailNowf("unable to close database", err.Error())
	}

	its.db, err = gorm.Open(postgres.Open(connection), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		its.FailNowf("unable to connect to database", err.Error())
	}

	tx = its.db.Exec(dropaDatabase)
	if tx.Error != nil {
		its.FailNowf("unable to drop database", tx.Error.Error())
	}

}
