package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type config struct {
	Postgres struct {
		DBMS   string
		USER   string
		PASS   string
		HOST   string
		PORT   string
		DBNAME string
		OPTION string
	}
}

// C is Config Valiable
var C config

// LoadConfig is Loading Application Config
func LoadConfig() {
	Conf := &C

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join("config"))

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("[Error] Load Config Error!")
		fmt.Println(err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		fmt.Println("[Error] Config Unmarshal Error!")
		fmt.Println(err)
		os.Exit(1)
	}
}
