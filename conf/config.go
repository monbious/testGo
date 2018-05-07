package conf

import (
	"testgopb/conf"
)

type Configs struct {
	conf.ConfigServiceMicro
}

var Config Configs

func init() {
	conf.LoadConf(&Config.ConfigServiceMicro,"conf/microservice.yaml")
}

