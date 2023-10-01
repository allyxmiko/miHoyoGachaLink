package utils

import (
	"io"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
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
func RegexpParser(pattern string) *regexp.Regexp {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatalln("正则表达式解析失败！")
	}
	return reg
}

func compareVersions(version1, version2 string) int {
	parts1 := strings.Split(version1, ".")
	parts2 := strings.Split(version2, ".")

	// 比较版本号的每个部分
	for i := 0; i < len(parts1) && i < len(parts2); i++ {
		part1 := parts1[i]
		part2 := parts2[i]

		// 将部分转换为整数进行比较
		num1 := atoi(part1)
		num2 := atoi(part2)

		if num1 < num2 {
			return -1
		} else if num1 > num2 {
			return 1
		}
	}

	// 如果前面的部分都相同，较长的版本号更大
	if len(parts1) < len(parts2) {
		return -1
	} else if len(parts1) > len(parts2) {
		return 1
	}

	// 版本号相同
	return 0
}

func atoi(s string) int {
	num := 0
	for i := 0; i < len(s); i++ {
		num = num*10 + int(s[i]-'0')
	}
	return num
}

func FindMaxVersion(versions []string) string {
	if len(versions) == 0 {
		return ""
	}

	maxVersion := versions[0]
	for i := 1; i < len(versions); i++ {
		if compareVersions(versions[i], maxVersion) > 0 {
			maxVersion = versions[i]
		}
	}

	return maxVersion
}
