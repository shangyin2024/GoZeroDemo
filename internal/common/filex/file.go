package filex

import (
	"context"
	"io"
	"net/http"
	"os"

	"github.com/zeromicro/go-zero/core/logx"
)

func GetFileName(ctx context.Context, path string) []byte {
	b, err := os.ReadFile(path)
	if err != nil {
		// logx.WithContext(ctx).Errorf("read file error, file path: %v", path)
		return nil
	}
	return b
}

// DownloadFile 下载文件本保存到本地
func DownloadFile(ctx context.Context, url, fPath string) error {
	r, err := http.Get(url)
	if err != nil {
		logx.WithContext(ctx).Errorf("download file error, err: %v", err)
		return err
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		logx.WithContext(ctx).Errorf("read file error, err: %v", err)
		return err
	}

	// 保存文件
	err = os.WriteFile(fPath, b, 066)
	if err != nil {
		logx.WithContext(ctx).Errorf("write file error, err: %v", err)
		return err
	}

	return nil
}

func DeleteFile(ctx context.Context, path string) error {
	err := os.Remove(path)
	if err != nil {
		logx.WithContext(ctx).Errorf("delete file error, err: %v", err)
		return err
	}
	return nil
}

// CreateIfNotExists 判断文件夹是否存在，存在则之际返回，不存在则创建文件夹
func CreateIfNotExists(ctx context.Context, directoryPath string) error {
	// 检查文件夹是否存在
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		// 文件夹不存在，创建它
		err := os.MkdirAll(directoryPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
