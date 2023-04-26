package mysql

import (
	"fmt"
	"log"
	"os"
	"time"

	// driver mysql on this implementation
	_ "github.com/go-sql-driver/mysql"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

type infoDatabaseMySQL struct {
	Read struct {
		Hostname   string
		Name       string
		Username   string
		Password   string
		Port       string
		Parameter  string
		Timezone   string
		DriverConn string
	}
	Write struct {
		Hostname   string
		Name       string
		Username   string
		Password   string
		Port       string
		Parameter  string
		Timezone   string
		DriverConn string
	}
}

func (infoDB *infoDatabaseMySQL) getMysqlConn(nameMap string) (err error) {

	viper.SetConfigFile("config.json")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = mapstructure.Decode(viper.GetStringMap(nameMap), infoDB)
	if err != nil {
		return
	}

	infoDB.Read.DriverConn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		infoDB.Read.Username, infoDB.Read.Password, infoDB.Read.Hostname, infoDB.Read.Port, infoDB.Read.Name)
	infoDB.Write.DriverConn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		infoDB.Write.Username, infoDB.Write.Password, infoDB.Write.Hostname, infoDB.Write.Port, infoDB.Write.Name)
	return
}

func initMysqlDB(inGormDB *gorm.DB, infoPg infoDatabaseMySQL) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	inGormDB, err := gorm.Open(mysql.Open(infoPg.Write.DriverConn), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	err = inGormDB.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{mysql.Open((infoPg.Read.DriverConn))},
	}))
	if err != nil {
		return nil, err
	}

	return inGormDB, nil
}
