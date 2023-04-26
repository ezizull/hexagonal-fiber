package jwt

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

func getTimeExpire(tokenType string) (tokenTimeUnix time.Duration, err error) {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		_ = fmt.Errorf("fatal error in config file: %s", err.Error())
	}

	JWTExpTime := viper.GetString(TokenTypeExpTime[tokenType])
	tokenTimeConverted, err := strconv.ParseInt(JWTExpTime, 10, 64)
	if err != nil {
		return
	}

	tokenTimeUnix = time.Duration(tokenTimeConverted)
	switch tokenType {
	case Refresh:
		tokenTimeUnix *= time.Hour
	case Access:
		tokenTimeUnix *= time.Minute
	default:
		err = errors.New("invalid token type")
		return
	}

	return
}
