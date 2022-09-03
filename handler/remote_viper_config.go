package handler

import (
	"example.com/http_demo/config"
	"fmt"
	"net/http"
)

func RemoteViperConfig(w http.ResponseWriter, r *http.Request) {
	// 此例子，需要启用远程配置中心后才能使用
	fmt.Fprintf(w, "Redis field %s, Value %s\n", "address", config.Redis.Address)

}
