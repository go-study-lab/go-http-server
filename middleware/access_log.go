package middleware

import (
	"bytes"
	"example.com/http_demo/utils/vlog"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type ResponseWithRecorder struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (rec *ResponseWithRecorder) WriteHeader(statusCode int) {
	rec.ResponseWriter.WriteHeader(statusCode)
	rec.statusCode = statusCode
}

func (rec *ResponseWithRecorder) Write(d []byte) (n int, err error) {
	n, err = rec.ResponseWriter.Write(d)
	if err != nil {
		return
	}
	rec.body.Write(d)

	return
}

// 创建中间件

func AccessLogging(f http.Handler) http.Handler {

	// 创建一个新的handler包装http.HandlerFunc
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		r.Body.Close()
		r.Body = io.NopCloser(buf)
		logEntry := vlog.AccessLog.WithFields(logrus.Fields{
			"ip":           r.RemoteAddr,
			"method":       r.Method,
			"path":         r.RequestURI,
			"query":        r.URL.RawQuery,
			"request_body": buf.String(),
		})

		wc := &ResponseWithRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
			body:           bytes.Buffer{},
		}

		// 调用下一个中间件或者最终的handler处理程序
		f.ServeHTTP(wc, r)

		defer logEntry.WithFields(logrus.Fields{
			"status":        wc.statusCode,
			"response_body": wc.body.String(),
		}).Info()

	})
}
