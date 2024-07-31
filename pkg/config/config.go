package config

import "github.com/spf13/viper"

type Config struct {
	DB_URL    string `mapstructure:"DB_URL"`
}

func LoadConfig() (err error) {
	var config Config
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&config)
	return err

}
