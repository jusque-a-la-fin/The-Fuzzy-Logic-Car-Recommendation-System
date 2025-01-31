package config

import "github.com/spf13/viper"

func SetupConfig(name string) error {
	viper.AddConfigPath("../../configs")
	setup(name)
	return viper.ReadInConfig()
}

func SetupConfigForService(name string) error {
	viper.AddConfigPath("../../../configs/services")
	setup(name)
	return viper.ReadInConfig()
}

func setup(name string) {
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
}
