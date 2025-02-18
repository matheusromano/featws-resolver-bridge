package config

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//Config ...
type Config struct {
	Port          string `mapstructure:"PORT"`
	ResolversFile string `mapstructure:"FEATWS_RESOLVER_BRIDGE_RESOLVERS_FILE"`
}

var config = &Config{}

//LoadConfig ...
func LoadConfig() (err error) {
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.SetDefault("PORT", "9000")
	viper.SetDefault("FEATWS_RESOLVER_BRIDGE_RESOLVERS_FILE", "./resolvers.json")

	err = viper.ReadInConfig()
	if err != nil {
		if err2, ok := err.(*os.PathError); !ok {
			err = err2
			log.Errorf("Error on Load Config: %v", err)
			return
		}
	}

	err = viper.Unmarshal(config)

	return
}

// GetConfig ...
func GetConfig() *Config {
	return config
}
