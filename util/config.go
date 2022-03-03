package util

import "github.com/spf13/viper"

type Config struct {
	PORT string `yaml:"PORT"`
	DBDriver string `yaml:"DB_DRIVER"`
	DBSource string `yaml:"DB_SOURCE"`
	RPCURL	string 	`yaml:"RPCURL"`
}


func LoadConfig(path string) (c Config, err error){
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&c)
	return
}
