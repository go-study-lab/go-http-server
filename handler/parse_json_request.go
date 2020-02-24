package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func ParseJsonRequestHandler(w http.ResponseWriter, r *http.Request) {
	var p Person

	// 将请求体中的 JSON 数据解析到结构体中
	// 发生错误，返回400 错误码
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Person: %+v", p)
}
