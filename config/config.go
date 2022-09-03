package config

import (
	"github.com/spf13/viper"
	"time"
)

type databaseConfig struct {
	Type        string        `mapstructure:"type"`
	DSN         string        `mapstructure:"dsn"`
	MaxOpenConn int           `mapstructure:"maxopen""`
	MaxIdleConn int           `mapstructure:"maxidle"`
	MaxLifeTime time.Duration `mapstructure:"maxlifetime"`
}

var Database *databaseConfig

func init() {
	vp := viper.New()
	vp.AddConfigPath("config/")
	// 产品级项目，可以设置成根据环境环境加载配置目录
	// env := os.Getenv("ENV")
	// vp.AddConfigPath("config/" + env)
	// 不设置配置名的时候，则读取目录下的所有文件
	// vp.SetConfigName("config")
	// 指定只读取 yaml 型的配置文件
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		panic(err)
	}
	vp.UnmarshalKey("database", &Database)
	Database.MaxLifeTime *= time.Second
}
