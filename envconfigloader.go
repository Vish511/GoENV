package config

import (
	"fmt"
	"os"
	"reflect"
)

/* Tag names to load configuration from environment variable */
const (
	ENV     = "env"
	DEFAULT = "default"
)

type Configuration struct {
	Port        string `env:"port" default:"3009"`
	MongoURL    string `env:"MongoUrl" default:"mongodb://localhost:27017/test"`
	UserService string `env:"UserService" default:"http://localhost:3005"`
	AuthService string `env:"AuthService" default:"http://localhost:3050"`
	Debug       string `env:"Debug" default:"true"`
}

/* Non-exported instance to avoid accidental overwrite */
var serviceConfig Configuration

func setConfig() {
	// ValueOf returns a Value representing the run-time data
	v := reflect.ValueOf(serviceConfig)
	for i := 0; i < v.NumField(); i++ {
		// Get the field tag value
		tag := v.Type().Field(i).Tag.Get(ENV)
		defaultTag := v.Type().Field(i).Tag.Get(DEFAULT)

		// Skip if tag is not defined or ignored
		if tag == "" || tag == "-" {
			continue
		}
		a := reflect.Indirect(reflect.ValueOf(serviceConfig))
		EnvVar, Info := loadFromEnv(tag, defaultTag)
		if Info != "" {
			fmt.Println("Missing environment configuration for '" + a.Type().Field(i).Name + "', Loading default setting!")
		}
		/* Set the value in the environment variable to the respective struct field */
		reflect.ValueOf(&serviceConfig).Elem().Field(i).SetString(EnvVar)
	}
}

func loadFromEnv(tag string, defaultTag string) (string, string) {
	/* Check if the tag is defined in the environment or else replace with default value */
	envVar := os.Getenv(tag)
	if envVar == "" {
		envVar = defaultTag
		/* '1' is used to indicate that default value is being loaded */
		return envVar, "1"
	}
	return envVar, ""
}

/*GetConfiguration :Exported function to return a copy of the configuration instance */
func GetConfiguration() Configuration {
	return serviceConfig
}

func init() {
	setConfig()
	fmt.Printf("Service configuration : %+v\n ", serviceConfig)
}
