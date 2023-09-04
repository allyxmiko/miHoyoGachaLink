package gacha

import (
	"bufio"
	"errors"
	"io"
	"log"
	"miHoyoGachaLink/config"
	"miHoyoGachaLink/utils"
	"os"
	"path/filepath"
	"time"
)

type Gacha struct {
	GachaLink   string
	GachaType   string
	LogPath     string
	DataPath    string
	DataPartten string
	LinkPartten string
}

func NewGacha(gachaType string) *Gacha {

	gachaConfig := config.NewGachaConfig(gachaType)
	gacha := &Gacha{
		GachaType:   gachaType,
		DataPartten: gachaConfig.DataPartten,
		LinkPartten: gachaConfig.LinkPartten,
	}
	gacha.GetLogPath(gachaConfig.BaseLogPath)
	gacha.GetDataPath(gachaConfig.BaseDataPath)
	return gacha
}

func (g *Gacha) GetLogPath(baseLog string) {
	logPath := filepath.Join(utils.HomeDir(), baseLog)
	if file, err := os.Stat(logPath); err != nil {
		log.Fatalln(file.Name() + "文件不存在！")
	}
	g.LogPath = utils.CopyFileToTemp(logPath)
}

func (g *Gacha) GetDataPath(baseData string) {
	var file *os.File
	var err error
	if file, err = os.OpenFile(g.LogPath, os.O_RDONLY, 0666); err != nil {
		log.Fatalln("日志文件读取失败！")
	}
	defer func(file *os.File) {
		if err = file.Close(); err != nil {
			log.Fatalln(err)
		}
	}(file)

	reader := bufio.NewReader(file)
	var match string
	var line []byte
	for {
		if line, _, err = reader.ReadLine(); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
		}
		list := utils.RegexpParser(g.DataPartten, string(line))
		if len(list) != 0 {
			match = list[0][0]
			break
		}
	}

	if match == "" {
		log.Fatalln("未解析到数据文件路径！")
	}

	dataPath := filepath.Join(match, baseData)
	var fi os.FileInfo
	if fi, err = os.Stat(dataPath); err != nil {
		log.Fatalln(fi.Name() + "文件不存在！")
	}
	expTime := int64(2 * 3600)
	modTime := fi.ModTime().Unix()
	now := time.Now().Unix()
	if now-modTime > expTime {
		log.Fatalln("数据文件已过期！请登录" + g.GachaType + "查看抽卡记录以刷新有效期！")
	}
	g.DataPath = utils.CopyFileToTemp(dataPath)
}

func (g *Gacha) ParseGachaLink() string {
	var data []byte
	var err error
	if data, err = os.ReadFile(g.DataPath); err != nil {
		log.Fatalln("文件读取失败！")
	}
	match := utils.RegexpParser(g.LinkPartten, string(data))
	if len(match) == 0 {
		log.Fatalln("未匹配到抽卡链接!请登录" + g.GachaType + "查看抽卡记录以刷新链接！")
	}
	return match[len(match)-1][0]
}

func (g *Gacha) GetGachaLink() string {
	// 返回链接
	link := g.ParseGachaLink()
	// 清理临时文件
	utils.CleanTempFile(g.LogPath, g.DataPath)
	return link
}
