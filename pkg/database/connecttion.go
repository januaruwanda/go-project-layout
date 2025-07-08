package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var connection *gorm.DB

func GetConnection() *gorm.DB {
	return connection
}

func NewConnection() {
	host := viper.Get("database.host")
	port := viper.Get("database.port")
	user := viper.Get("database.user")
	password := viper.Get("database.password")
	databaseName := viper.Get("database.name")

	destination := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?parseTime=true",
		user,
		password,
		host,
		port,
		databaseName,
	)

	database, error := gorm.Open(mysql.Open(destination), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if error != nil {
		panic(error)
	}

	sqlDatabase, error := database.DB()
	if error != nil {
		panic(error)
	}

	sqlDatabase.SetMaxIdleConns(1)
	sqlDatabase.SetMaxOpenConns(1)
	sqlDatabase.SetConnMaxLifetime(time.Hour)
	connection = database
}

type DatabaseHandlerFunc func(dest interface{}, isExec bool, query string, values ...interface{}) error

func CreateDatabaseHandler() func(dest interface{}, isExecute bool, query string, values ...interface{}) error {

	if viper.GetBool(`debug`) {
		return func(dest interface{}, isExecute bool, query string, values ...interface{}) error {
			db := GetConnection()
			if isExecute {
				result := db.Exec(query, values...)
				if result.Error != nil {
					log.Error().Err(result.Error).Msg("error connection database")
				}
				return result.Error
			}
			return db.Raw(query, values...).Scan(dest).Error
		}
	} else {
		return func(dest interface{}, isExecute bool, query string, values ...interface{}) error {
			db := GetConnection()
			if isExecute {
				result := db.Debug().Exec(query, values...)
				if result.RowsAffected == 0 {
					log.Error().Err(result.Error).Msg("no rows were inserted")
					return errors.New("no rows were inserted")
				}
				return result.Error
			}
			return db.Debug().Raw(query, values...).Scan(dest).Error
		}

	}

}
