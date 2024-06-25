package middleware

import (
	"errors"
	"net/http"
	"strconv"

	xhttp "github.com/zeromicro/x/http"
)

type ParameterValidationMiddleware struct {
}

func NewParameterValidationMiddleware() *ParameterValidationMiddleware {
	return &ParameterValidationMiddleware{}
}

// Handle parameter validation middleware
func (m *ParameterValidationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		pageStr := values.Get("page")
		if pageStr != "" {
			pageInt, _ := strconv.Atoi(pageStr)
			if pageInt > 10000 {
				xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New("page must less than 10000"))
				return
			}
		}
		pageSizeStr := values.Get("page_size")
		if pageSizeStr != "" {
			pageSizeInt, _ := strconv.Atoi(pageSizeStr)
			if pageSizeInt > 1000 {
				xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New("page size must less than 1000"))
				return
			}
		}
		next(w, r)
	}
}
