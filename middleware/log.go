package middleware

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// 记录每个URL请求的执行时长
func Logging() mux.MiddlewareFunc {

	// 创建中间件
	return func(f http.Handler) http.Handler {

		// 创建一个新的handler包装http.HandlerFunc
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// 中间件的处理逻辑
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			// 调用下一个中间件或者最终的handler处理程序
			f.ServeHTTP(w, r)
		})
	}
}
