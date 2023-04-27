package redis

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type infoDatabaseRedis struct {
	Read struct {
		Hostname   string
		Username   string
		Password   string
		Port       string
		Database   string
		DriverConn string
	}
	Write struct {
		Hostname   string
		Username   string
		Password   string
		Port       string
		Database   string
		DriverConn string
	}
}

// Database cradential
var (
	hostname = os.Getenv("REDIS_HOST")
	port     = os.Getenv("REDIS_PORT")
	username = os.Getenv("REDIS_USER")
	password = os.Getenv("REDIS_PASSWORD")
	dbname   = os.Getenv("REDIS_DBNAME")
	dbtime   = os.Getenv("REDIS_DBTIME")
)

func (infoDB *infoDatabaseRedis) getRedisConn(nameMap string) (err error) {

	viper.SetConfigFile("config.json")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = mapstructure.Decode(viper.GetStringMap(nameMap), infoDB)
	if err != nil {
		return
	}

	if hostname != "" {
		infoDB.Read.Hostname = hostname
		infoDB.Write.Hostname = hostname
	}

	if port != "" {
		infoDB.Read.Port = port
		infoDB.Write.Port = port
	}
	if username != "" {
		infoDB.Read.Username = username
		infoDB.Write.Username = username
	}
	if password != "" {
		infoDB.Read.Password = password
		infoDB.Write.Password = password
	}

	if dbname != "" {
		infoDB.Read.Database = dbname
		infoDB.Write.Database = dbname
	}

	infoDB.Read.DriverConn = fmt.Sprintf("redis://%s:%s@%s:%s/%s",
		infoDB.Read.Username, infoDB.Read.Password, infoDB.Read.Hostname, infoDB.Read.Port, infoDB.Read.Database)
	infoDB.Write.DriverConn = fmt.Sprintf("redis://%s:%s@%s:%s/%s",
		infoDB.Write.Username, infoDB.Write.Password, infoDB.Write.Hostname, infoDB.Write.Port, infoDB.Write.Database)

	return
}

func initRedisDB(inRedisDB *redis.Client, infoRed infoDatabaseRedis) (*redis.Client, error) {
	database, err := strconv.ParseInt(infoRed.Write.Database, 10, 64)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "error when connect to repository")
	}

	inRedisDB = redis.NewClient(&redis.Options{
		Addr:     infoRed.Write.Hostname + ":" + infoRed.Write.Port,
		Username: infoRed.Write.Username,
		Password: infoRed.Write.Password,
		DB:       int(database),
	})

	return inRedisDB, nil
}
