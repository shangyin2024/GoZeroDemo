package filex

import "strings"

// imageContentTypeMap 是一个仅包含常见图片Content-Type到文件扩展名的映射
var imageContentTypeMap = map[string]string{
	"image/jpeg":    ".jpg",
	"image/png":     ".png",
	"image/gif":     ".gif",
	"image/svg+xml": ".svg",
	"image/webp":    ".webp",
	"image/bmp":     ".bmp",
	"image/tiff":    ".tiff",
	"image/x-icon":  ".ico",
}

// GetImageExtFromContentType 根据图片的Content-Type获取文件扩展名
func GetImageExtFromContentType(contentType string) string {
	ext, exists := imageContentTypeMap[strings.ToLower(contentType)]
	if exists {
		return ext
	}
	return "" // 如果没有找到对应的类型，则返回空字符串
}
