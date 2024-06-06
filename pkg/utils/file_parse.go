package utils

import (
	"bytes"
	"io"
	"mime/multipart"
)

// FileHeaderToBytes 将 *multipart.FileHeader 转换为 []byte
func FileHeaderToBytes(fileHeader *multipart.FileHeader) ([]byte, error) {
	// 打开文件
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 创建一个缓冲区
	var buf bytes.Buffer

	// 将文件内容拷贝到缓冲区
	_, err = io.Copy(&buf, file)
	if err != nil {
		return nil, err
	}

	// 返回文件内容的 []byte 表示
	return buf.Bytes(), nil
}
