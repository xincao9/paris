package config

import (
	"github.com/spf13/viper"
	"log"
	"paris/meta"
)

var (
	C *viper.Viper
)

func init() {
	// 配置文件设置
	C = viper.New()
	C.SetConfigName("config")
	C.SetConfigType("yaml")
	C.AddConfigPath("/etc/paris/")
	C.AddConfigPath("$HOME/.paris")
	C.AddConfigPath(".")
	C.SetDefault("data.dir", meta.DefaultDir)
	C.SetDefault("server.mode", "debug")
	C.SetDefault("server.port", ":9090")
	C.SetDefault("logger.level", "debug")
	err := C.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config : %v\n", err)
	}
}
