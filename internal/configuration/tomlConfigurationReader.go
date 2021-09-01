package configuration

import (
	"github.com/BurntSushi/toml"
)

type TomlConfigurationReader struct{}

func (r *TomlConfigurationReader) read() (*GeneralConfiguration, error) {
	conf := GeneralConfiguration{}
	_, err := toml.DecodeFile("./configuration/configuration.toml", &conf)

	if err != nil {
		return nil, err
	}

	return &conf, nil
}
