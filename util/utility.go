package util

import (
	"github.com/spf13/viper"
	"log"
)

type ProductStock struct {
	Product  string
	Quantity int
}

type Config struct {
	Name               string
	BasicAuthorization string
	Ports              []Port
	Stock              []ProductStock
	CartStore          string
	CartDBConfig       []DBconfig
	StockStore         string
	StockDBConfig      []DBconfig
}

type Port struct {
	Key  string
	Port string
}

type DBconfig struct {
	Key string
	URL string
}

func ReadConfig() *Config {
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	var config Config

	_ = viper.Unmarshal(&config)

	return &config
}
