package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Event string

const (
	TurnOn  Event = "turnon"
	TurnOff Event = "turnoff"
)

type SwitchEvents struct {
	On  Event `mapstructure:"on"`
	Off Event `mapstructure:"off"`
}

type Switch struct {
	Name   string       `mapstructure:"name"`
	Events SwitchEvents `mapstructure:"events"`
}

type Config struct {
	Key      string   `mapstructure:"key"`
	Switches []Switch `mapstructure:"switches"`
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	viper.AddConfigPath(home + "/.govimar")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	for _, cf := range StaticConfigFields {
		if !viper.InConfig(string(cf)) {
			log.Fatalf("missing configuration field: %s", cf)
		}
		if !viper.IsSet(string(cf)) {
			log.Fatalf("void configuration field: %s", cf)
		}
		if cf == Switches {
			var sws []Switch
			err := viper.UnmarshalKey(string(cf), &sws)
			if err != nil {
				log.Fatal(err)
			}
			for i, sw := range sws {
				if sw.Name == "" {
					log.Fatalf("void configuration field: %s[%d].name", cf, i)
				}
				if sw.Events.On == "" {
					log.Fatalf("void configuration field: %s[%d].events.on", cf, i)
				}
				if sw.Events.Off == "" {
					log.Fatalf("void configuration field: %s[%d].events.off", cf, i)
				}
			}
		}
	}
}

// GetConfig returns the configuration for the application.
func GetConfig() *Config {
	c := &Config{}
	err := viper.Unmarshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
