package database

import (
	"hexagonal-fiber/infrastructure/repository/redis"

	"gorm.io/gorm"
)

type Database struct {
	Postgre *gorm.DB
	Redis   *redis.InfoDatabaseRedis
}
