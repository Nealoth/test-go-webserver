package configuration

import (
	"log"
)

var cm *configurationManager

type configurationManager struct {
	cachedConfiguration *GeneralConfiguration
}

func InitConfigurationManager() {
	cm = &configurationManager{}
}

func ConfigurationManager() *configurationManager {
	return cm
}

func (cm *configurationManager) GetAndCache(reader Reader) *GeneralConfiguration {
	conf, err := ReadConfiguration(reader)

	if err != nil {
		//TODO log
		log.Fatal(err)
	}

	cm.cachedConfiguration = conf
	return conf
}

func (cm *configurationManager) GetCached() *GeneralConfiguration {
	//Todo check nil
	return cm.cachedConfiguration
}
