package redis

import "github.com/redis/go-redis/v9"

// NewRedis is a function that returns a redis database connection using  initial configuration
func NewRedis() (redisDB *redis.Client, err error) {
	var infoPg infoDatabaseRedis
	err = infoPg.getRedisConn("Databases.Redis.Localhost")
	if err != nil {
		return nil, err
	}

	redisDB, err = initRedisDB(redisDB, infoPg)
	if err != nil {
		return nil, err
	}

	return
}
