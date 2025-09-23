package config

import (
	"context"
	"log/slog"
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Firebase Firebase `mapstructure:"FIREBASE"`
	Database    struct {
		Master Database `mapstructure:"MASTER"`
		Slave  Database `mapstructure:"SLAVE"`
	} `mapstructure:"DATABASE"`
}

type Firebase struct {
	CredentialFile string `mapstructure:"CREDENTIAL_FILE"`
}

type Database struct {
	Host     string `mapstructure:"HOST"`
	Port     int    `mapstructure:"PORT"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	Database string `mapstructure:"DATABASE"`
}


var (
	config *Config
	once   sync.Once
)

func GetConfig(ctx context.Context) *Config {
	once.Do(getConfig)
	return config
}

func getConfig() {
	ctx := context.TODO()
	// viper.AllowEmptyEnv(false)
	viper.SetConfigName("config")
	viper.AddConfigPath(".")    //Viper尋找設定檔的路徑，一次可以設定多個路徑
	viper.SetConfigType("yaml") //設定檔的檔案類型i
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		slog.ErrorContext(ctx, "error reading config file", "error", err)
		os.Exit(1)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		slog.ErrorContext(ctx, "unmarshal config file", "error", err)
		os.Exit(1)
	}
}
