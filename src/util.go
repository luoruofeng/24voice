package src

import (
	"os"
	"path/filepath"
)

func FileToBytes(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetFileExtensionWithoutDot 获取文件路径的文件类型后缀并去掉点
func GetFileExtensionWithoutDot(filePath string) string {
	ext := filepath.Ext(filePath)
	if ext != "" {
		// 如果文件有扩展名，去掉扩展名中的点号
		return ext[1:]
	}
	return "" // 文件没有扩展名
}
