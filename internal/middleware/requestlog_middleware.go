package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"slices"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type RequestLogMiddleware struct {
}

func NewRequestLogMiddleware() *RequestLogMiddleware {
	return &RequestLogMiddleware{}
}

// RequestLogMiddleware request log middleware
func (m *RequestLogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		method := r.Method

		var bodyBytes []byte
		if slices.Contains([]string{"GET", "POST", "PUT", "DELETE"}, method) {
			bodyBytes, err = io.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err)
			}
			// logx.WithContext(r.Context()).Infof("req: %v", string(bodyBytes))
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// 创建一个自定义的 ResponseWriter，用于记录响应
		// recorder := &responseRecorder{
		// 	ResponseWriter: w,
		// 	statusCode:     http.StatusOK,
		// 	body:           make([]byte, 0),
		// }
		// next(recorder, r)

		// responseBoy := string(recorder.body)

		t := time.Since(startTime).Seconds()

		logx.WithContext(r.Context()).Info(
			"ip:", httpx.GetRemoteAddr(r),
			"path", r.URL.Path,
			"method:", method,
			"query:", r.URL.Query().Encode(),
			"body", string(bodyBytes),
			"latency:", fmt.Sprintf("%.3f", t),
		)

		// Passthrough to next handler if need
		next(w, r)
	}
}

// 自定义的 ResponseWriter
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

// WriteHeader 重写 WriteHeader 方法，捕获状态码
func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

// 重写 Write 方法，捕获响应数据
func (r *responseRecorder) Write(body []byte) (int, error) {
	r.body = body
	return r.ResponseWriter.Write(body)
}
