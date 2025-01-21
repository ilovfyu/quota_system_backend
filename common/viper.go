package common

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	g "quota_system/global"
	"quota_system/model"
)

// 配置初始化
func InitViper() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Errorf("Fatal error config file: %s \n", err)
		panic(err)
	}
	var config *model.BaseConfig
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Errorf("Fatal error config file: %s \n", err)
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(&config); err != nil {
			fmt.Errorf("Fatal error config file: %s \n", err)
			panic(err)
		}
	})
	g.ServiceConfig = config
}
