package resource

import (
	"github.com/rmukubvu/amakhosi/model"
	"github.com/spf13/viper"
)

const (
	DbUserNameKey = "database.user"
	DbPwdKey      = "database.pwd"
	DbHostKey     = "database.host"
	DbPortKey     = "database.port"
)

var dbConfig model.DatabaseConfig

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		//we cant continue if the resource is missing as we
		//wont be able to connect to the database
		panic(err.Error())
	}

	dbConfig.User = viper.GetString(DbUserNameKey)
	dbConfig.Pwd = viper.GetString(DbPwdKey)
	dbConfig.Host = viper.GetString(DbHostKey)
	dbConfig.Port = viper.GetInt(DbPortKey)
}

func DbConfig() model.DatabaseConfig {
	return dbConfig
}
