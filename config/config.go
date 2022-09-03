package config

import (
	"github.com/spf13/viper"
	"path"
	"runtime"
	"time"
)

type databaseConfig struct {
	Type        string        `mapstructure:"type"`
	DSN         string        `mapstructure:"dsn"`
	MaxOpenConn int           `mapstructure:"maxopen""`
	MaxIdleConn int           `mapstructure:"maxidle"`
	MaxLifeTime time.Duration `mapstructure:"maxlifetime"`
}

type RedisConfig struct {
	Address   string `mapstructure:"address"`
	Password  string `mapstructure:"password"`
	DbNumber  int    `mapstructure:"dbnumber"`
	MaxActive int    `mapstructure:"maxactive"`
	MaxIdle   int    `mapstructure:"maxidle"`
}

var Database *databaseConfig
var Redis *RedisConfig

func init() {
	// 获取当前文件的路径
	_, filename, _, _ := runtime.Caller(0)
	// 配置文件目录的路径
	configBaseDir := path.Dir(filename)
	vp := viper.New()
	vp.AddConfigPath(configBaseDir)
	// 产品级项目，可以设置成根据环境环境加载配置目录
	// env := os.Getenv("ENV")
	// vp.AddConfigPath(configBaseDir + env)
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
	// 此配置没有值，计划使用远程配置中心配置
	// 请看远程配置中心版本 config_remote_version.go 的代码
	vp.UnmarshalKey("redis", &Redis)
}
