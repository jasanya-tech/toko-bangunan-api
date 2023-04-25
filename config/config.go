package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
)

var (
	cfg    Config
	doOnce sync.Once
)

type Config struct {
	Application struct {
		Port string `mapstructure:"PORT"`
	} `mapstructure:"APPLICATION"`

	Auth struct {
		JwtToken struct {
			DefaultToken string `mapstructure:"DEFAULT_KEY_TOKEN"`
			AccessToken  struct {
				PrivateKey string `mapstructure:"PRIVATE"`
				PublicKey  string `mapstructure:"PUBLIC"`
				Expired    int    `mapstructure:"EXPIRED"`
				MaxAge     int    `mapstructure:"MAX_AGE"`
			} `mapstructure:"ACCESS_TOKEN"`
			RefreshToken struct {
				PrivateKey string `mapstructure:"PRIVATE"`
				PublicKey  string `mapstructure:"PUBLIC"`
				Expired    int    `mapstructure:"EXPIRED"`
				MaxAge     int    `mapstructure:"MAX_AGE"`
			} `mapstructure:"REFRESH_TOKEN"`
		} `mapstructure:"JWT_TOKEN"`
	} `mapstructure:"AUTH"`

	DB struct {
		Mysql struct {
			Host string `mapstructure:"HOST"`
			Port int    `mapstructure:"PORT"`
			User string `mapstructure:"USER"`
			Pass string `mapstructure:"PASS"`
			Name string `mapstructure:"NAME"`
		} `mapstructure:"MYSQL"`
	} `mapstructure:"DB"`
}

func Get() Config {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf(fmt.Sprintf("cannot read .env file: %v", err))
	}
	doOnce.Do(func() {
		if err := viper.Unmarshal(&cfg); err != nil {
			log.Fatalln("cannot unmarshal config")
		}
	})
	return cfg
}
