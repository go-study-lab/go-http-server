package config

/**
* 如果使用远程配置中心，请用此源文件替换config.go
* 注意：需要有ETCD 环境才可以使用
 */

//import (
//	"example.com/http_demo/utils/zlog"
//	"github.com/spf13/viper"
//	_ "github.com/spf13/viper/remote"
//	"go.uber.org/zap"
//	"path"
//	"runtime"
//	"time"
//)
//
//type databaseConfig struct {
//	Type        string        `mapstructure:"type"`
//	DSN         string        `mapstructure:"dsn"`
//	MaxOpenConn int           `mapstructure:"maxopen""`
//	MaxIdleConn int           `mapstructure:"maxidle"`
//	MaxLifeTime time.Duration `mapstructure:"maxlifetime"`
//}
//
//type RedisConfig struct {
//	Address   string `mapstructure:"address"`
//	Password  string `mapstructure:"password"`
//	DbNumber  int    `mapstructure:"dbnumber"`
//	MaxActive int    `mapstructure:"maxactive"`
//	MaxIdle   int    `mapstructure:"maxidle"`
//}
//
//var Database *databaseConfig
//var Redis *RedisConfig
//
//func init() {
//	// 获取当前文件的路径
//	_, filename, _, _ := runtime.Caller(0)
//	// 配置文件目录的路径
//	configBaseDir := path.Dir(filename)
//	vp := viper.New()
//	vp.AddConfigPath(configBaseDir)
//	// 本地配置文件可以和远程配置中心一起使用
//	viper.RemoteConfig = &RemoteConfig{}
//	err := vp.AddRemoteProvider("etcd", "http://127.0.0.1:32379", "root/config/viper-test/config")
//	if err != nil {
//		panic(err)
//	}
//	// 产品级项目，可以设置成根据环境环境加载配置目录
//	// env := os.Getenv("ENV")
//	// vp.AddConfigPath("config/" + env)
//	// 不设置配置名的时候，则读取目录下的所有文件
//	// vp.SetConfigName("config")
//	// 指定只读取 yaml 型的配置文件
//	vp.SetConfigType("yaml")
//	err = vp.ReadInConfig()
//	if err != nil {
//		panic(err)
//	}
//	err = vp.ReadRemoteConfig()
//	if err != nil {
//		panic(err)
//	}
//	vp.UnmarshalKey("database", &Database)
//	Database.MaxLifeTime *= time.Second
//	vp.UnmarshalKey("redis", &Redis)
//
//	go watchRemoteConfig(vp)
//}
//
//func watchRemoteConfig(vp *viper.Viper) {
//	for {
//		time.Sleep(5 * time.Second)
//		err := vp.WatchRemoteConfigOnChannel()
//		if err != nil {
//			zlog.Error("Read Config Server Error", zap.Error(err))
//			return
//		}
//		// 监控远程配置的变更
//		vp.UnmarshalKey("redis", &Redis)
//	}
//}
