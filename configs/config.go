package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	viper.SetDefault("api.port", "9050")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "3060") // Porta padrão do Firebird
	viper.SetDefault("database.user", "SYSDBA")
	viper.SetDefault("database.pass", "masterkey")
	viper.SetDefault("database.name", "C:\\Users\\Usuario\\Desktop\\golang-firebird\\api_fire.fdb")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("ini")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}
	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}
	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}
	return nil
}
func GetDBDSN() string {
	conf := GetDB()
	return fmt.Sprintf("firebirdsql://%s:%s@%s:%s/%s",
	conf.User,
	conf.Pass,
	conf.Host,
	conf.Port,
	conf.Database,)

}

func GetDB() DBConfig {
	return cfg.DB
}
func GetServerPort() string {
	return cfg.API.Port
}
