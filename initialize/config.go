package initialize

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"video/global"
)

func InitConfig() {
	initViperConfig()
}

// initViperConfig 初始化配置 > viper 配置包
func initViperConfig() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panicf("获取配置异常: %s", err)
	}
	// 映射配置文件
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		zap.S().Panicf("加载配置异常: %s", err)
	}
	// 动态监听配置
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&global.ServerConfig)
	})
}
