package config

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

var Conf Config

func init() {
	if _, err := toml.Decode("config.toml", &Conf); err != nil {
		logrus.WithField("error", err).Fatal("Unable to parse config")
	}
}
