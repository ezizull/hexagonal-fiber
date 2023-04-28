package redis

// InitRedis is a function that returns a redis database connection using initial configuration
func InitRedis() (*InfoDatabaseRedis, error) {
	var infoDB InfoDatabaseRedis
	err := infoDB.getRedisConn("Databases.Redis.Localhost")
	if err != nil {
		return nil, err
	}

	return &infoDB, nil
}
