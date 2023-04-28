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

	redisDB, err := infoDB.NewRedis(0)
	if err != nil {
		return nil, err
	}

	pingStr, err := redisDB.Ping(Ctx).Result()
	if err != nil {
		return nil, err
	}

	fmt.Println("ping redis ", pingStr)

	return &infoDB, nil
}
