package utils

import (
	"io"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
)

// HomeDir 获取家目录
func HomeDir() string {
	var u *user.User
	var err error
	if u, err = user.Current(); err != nil {
		log.Fatalln("未获取到家目录")
	}
	return u.HomeDir
}

// CopyFileToTemp 将文件复制到临时目录
func CopyFileToTemp(filePath string) string {
	var srcFile *os.File
	var target *os.File
	var err error
	if srcFile, err = os.Open(filePath); err != nil {
		log.Fatalln("文件打开失败！请确保文件存在！")
	}
	defer func(srcFile *os.File) {
		if err = srcFile.Close(); err != nil {
			log.Fatalln("文件流关闭失败！")
		}
	}(srcFile)
	fileInfo, _ := srcFile.Stat()
	targetPath := filepath.Join(os.TempDir(), fileInfo.Name())
	if target, err = os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY, 0755); err != nil {
		log.Fatalln("临时文件创建失败！")
	}
	defer func(target *os.File) {
		if err = target.Close(); err != nil {
			log.Fatalln("文件流关闭失败！")
		}
	}(target)
	if _, err = io.Copy(target, srcFile); err != nil {
		log.Fatalln("文件复制出错！")
	}
	return target.Name()
}

// CleanTempFile 删除临时文件
func CleanTempFile(filePaths ...string) {
	for _, filePath := range filePaths {
		if err := os.Remove(filePath); err != nil {
			log.Fatalln("临时文件删除失败!请在", os.TempDir(), "目录下手动删除"+filePath+"文件!")
		}
	}
}

// RegexpParser 正则匹配
func RegexpParser(pattern string, content string) [][]string {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatalln("正则表达式解析失败！")
	}
	return reg.FindAllStringSubmatch(content, -1)
}
