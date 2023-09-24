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
	Link        string
	GachaType   string
	LogPath     string
	DataPath    string
	DataPartten string
	LinkPartten string
}

func NewGacha(gachaType string) (*Gacha, error) {
	var err error
	var gachaConfig *config.GachaConfig
	if gachaConfig, err = config.NewGachaConfig(gachaType); err != nil {
		return nil, err
	}
	gacha := &Gacha{
		GachaType:   gachaType,
		DataPartten: gachaConfig.DataPartten,
		LinkPartten: gachaConfig.LinkPartten,
	}
	if err = gacha.GetLogPath(gachaConfig.BaseLogPath); err != nil {
		return nil, err
	}
	if err = gacha.GetDataPath(gachaConfig.BaseDataPath); err != nil {
		return nil, err
	}
	if err = gacha.GetGachaLink(); err != nil {
		return nil, err
	}
	return gacha, nil
}

func (g *Gacha) GetLogPath(baseLog string) error {
	logPath := filepath.Join(utils.HomeDir(), baseLog)
	if file, err := os.Stat(logPath); err != nil {
		return NewFileNotExistError(file.Name())
	}
	g.LogPath = utils.CopyFileToTemp(logPath)
	return nil
}

func (g *Gacha) GetDataPath(baseData string) error {
	var file *os.File
	var err error
	if file, err = os.OpenFile(g.LogPath, os.O_RDONLY, 0666); err != nil {
		return NewFileReadError(file.Name())
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
		return NewDataFilePathParseError(file.Name())
	}

	dataPath := filepath.Join(match, baseData)
	var fi os.FileInfo
	if fi, err = os.Stat(dataPath); err != nil {
		return NewFileNotExistError(fi.Name())
	}
	expTime := int64(2 * 3600)
	modTime := fi.ModTime().Unix()
	now := time.Now().Unix()
	if now-modTime > expTime {
		return NewDataFileExpiredError(fi.Name(), g.GachaType)
	}
	g.DataPath = utils.CopyFileToTemp(dataPath)
	return nil
}

func (g *Gacha) ParseGachaLink() (string, error) {
	var data []byte
	var err error
	if data, err = os.ReadFile(g.DataPath); err != nil {
		return "", NewFileReadError(g.DataPath)
	}
	match := utils.RegexpParser(g.LinkPartten, string(data))
	if len(match) == 0 {
		return "", NewLinkParseError(g.GachaType)
	}
	return match[len(match)-1][0], nil
}

func (g *Gacha) GetGachaLink() error {
	var err error
	// 返回链接
	if g.Link, err = g.ParseGachaLink(); err != nil {
		return err
	}
	g.CleanTempFile()
	return nil
}

func (g *Gacha) CleanTempFile() {
	utils.CleanTempFile(g.LogPath, g.DataPath)
}
