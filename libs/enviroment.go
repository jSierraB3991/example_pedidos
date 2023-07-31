package libs

import (
	"log"
	"os"

	config "github.com/spf13/viper"
)

type Enviroment struct {
	DatabaseURL  string
	DatabaseName string
}

func Configure(ConfigPath string, ConfigName string) *Enviroment {
	config.AddConfigPath(ConfigPath)
	config.SetConfigName(ConfigName)

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if config.GetString("database.url") == "" {
		log.Println("NO database url in the configuration file" + ConfigName)
		os.Exit(1)
	}
	if config.GetString("database.name") == "" {
		log.Println("No name datase in the configuration file" + ConfigName)
		os.Exit(1)
	}

	return &Enviroment{
		config.GetString("database.url"),
		config.GetString("database.name"),
	}
}
