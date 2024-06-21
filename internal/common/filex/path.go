package filex

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"net/url"
	"strings"
)

// Process 处理图片链接
func Process(ctx context.Context, targetUrl, sourceUrl string) string {
	if sourceUrl == "" {
		return sourceUrl
	}
	parse, err := url.Parse(sourceUrl)
	if err != nil {
		logx.WithContext(ctx).Error(err)
		return sourceUrl
	}
	if parse.Scheme == "" { // 不是合法的url
		return strings.Trim(targetUrl, "/") + "/" + strings.Trim(sourceUrl, "/")
	}
	return sourceUrl
}
