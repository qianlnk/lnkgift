/**
 * FileName:		config.go
 * Description:		read config file
 * Author:			Qianno.Xie
 * Email:			qianlnk@163.com
**/
package config

import (
	"fmt"
)

import (
	"github.com/spf13/viper"
	"sync"
)

var (
	once sync.Once
	conf *viper.Viper
)

func init() {
	once.Do(initilaizing)
}

func initilaizing() {
	fmt.Println("start config")
	conf = viper.New()
	conf.SetConfigName("config")
	conf.SetConfigType("toml")
	conf.AddConfigPath("config")

	err := conf.ReadInConfig()
	if err != nil {
		return
	}
}

func GetConfig() *viper.Viper {
	return conf
}
