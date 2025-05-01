package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Load() error {
	viper.SetConfigFile(".env");
	if err := viper.ReadInConfig(); err != nil {
		return err;
	}

	return nil;
}

func DBConnectionString() string {
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbName);
}
