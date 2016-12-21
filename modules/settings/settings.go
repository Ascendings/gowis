package settings

import "gopkg.in/ini.v1"

var (
	// Cfg - configuration settings
	Cfg *ini.File
)

func init() {
	// load the config file
	cfg, cfgErr := ini.InsensitiveLoad("./conf/app.ini")

	// check for errors while loading the configuration
	if cfgErr != nil {
		panic(cfgErr)
	}

	Cfg = cfg
}
