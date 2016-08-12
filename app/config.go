package app

import "gopkg.in/ini.v1"

// InitConfig - returns the finalized configuration object
func InitConfig() *ini.File {
	// load the config file
	cfg, cfgErr := ini.InsensitiveLoad("./app/app.ini")

	// check for errors while loading the configuration
	if cfgErr != nil {
		panic(cfgErr)
	}

	return cfg
}
