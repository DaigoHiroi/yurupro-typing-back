package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

// Config conf struct
type Config struct {
	Server struct {
		Port int
		Mock bool
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
	}
}

// Current runnnig configuration
var Current *Config

// NewConfig プロジェクトのコンフィグ設定をロードします.
func NewConfig(runServer bool) {
	var C Config
	Current = &C
	// viper.AddConfigPath("$GOPATH/app/conf/")

	viper.AddConfigPath("conf/")
	viper.SetConfigType("yml")

	if runServer {
		viper.SetConfigName("production")
	} else {
		viper.SetConfigName("local")
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal conf file error: %s", err))
	}

	if err := viper.Unmarshal(&C); err != nil {
		panic(fmt.Errorf("fatal conf file error: %s", err))
	}
	return
}
