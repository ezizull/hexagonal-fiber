package redis

import (
	"fmt"
)

// InitRedis is a function that returns a redis database connection using initial configuration
func InitRedis() (*InfoDatabaseRedis, error) {
	var infoDB InfoDatabaseRedis
	err := infoDB.getRedisConn("Databases.Redis.Localhost")
	if err != nil {
		return nil, err
	}

	redisDB := infoDB.NewRedis(0)
	pingStr, err := redisDB.Ping(infoDB.CTX).Result()
	if err != nil {
		return nil, err
	}

	fmt.Println("redis ping ", pingStr)
	return &infoDB, nil
}
