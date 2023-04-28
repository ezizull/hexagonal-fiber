package redis

import "github.com/gofiber/fiber/v2"

// InitRedis is a function that returns a redis database connection using initial configuration
func InitRedis() (infoRed *InfoDatabaseRedis, err error) {
	if err = infoRed.getRedisConn("Databases.Redis.Localhost"); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "error when initial repository")
	}

	return infoRed, nil
}
