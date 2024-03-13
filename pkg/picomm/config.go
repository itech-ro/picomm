package picomm

import (
	"log"

	"github.com/spf13/viper"
)

// Config represents the configuration
type Config interface {
	GetString(string) string
	GetInt(string) int
	GetInt64(string) int64
	GetBool(string) bool
	GetStringMapString(key string) map[string]string
}

// GetConfig returns the configuration
func GetConfig(filename string, path string) (Config, error) {
	viper.SetConfigName(filename)
	viper.AddConfigPath(path)
	viper.AddConfigPath(".")

	// HTTP
	viper.BindEnv("http.address", "HTTP_ADDRESS")
	viper.BindEnv("http.port", "HTTP_PORT")
	viper.BindEnv("http.timeout.read", "HTTP_TIMEOUT_READ")
	viper.BindEnv("http.timeout.write", "HTTP_TIMEOUT_WRITE")

	// ----- Default Values -----
	// HTTP
	viper.SetDefault("http.address", "0.0.0.0:8080")
	viper.SetDefault("http.timeout.read", 30)
	viper.SetDefault("http.timeout.write", 30)

	viper.AutomaticEnv()

	viper.ReadInConfig()

	configFileUsed := viper.ConfigFileUsed()

	if len(configFileUsed) == 0 {
		log.Println("no configuration file found")
	} else {
		log.Printf("configuration file »%s« used\n", configFileUsed)
	}

	return viper.GetViper(), nil
}
