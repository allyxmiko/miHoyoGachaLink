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
	Link      string
	GachaType string
	Config    *config.GachaConfig
}

func NewGacha(gachaType string) (*Gacha, error) {
	var err error
	var gachaConfig *config.GachaConfig
	if gachaConfig, err = config.NewGachaConfig(gachaType); err != nil {
		return nil, err
	}
	gacha := &Gacha{
		GachaType: gachaType,
		Config:    gachaConfig,
	}
	if err = gacha.GetGachaLink(); err != nil {
		return nil, err
	}
	return gacha, nil
}

func (g *Gacha) GetLogPath(baseLog string) (string, error) {
	logPath := filepath.Join(utils.HomeDir(), baseLog)
	if file, err := os.Stat(logPath); err != nil {
		return "", NewFileNotExistError(file.Name())
	}
	return utils.CopyFileToTemp(logPath), nil
}

func (g *Gacha) GetDataPath(logPath, baseData string) (string, error) {
	var file *os.File
	var err error
	if file, err = os.OpenFile(logPath, os.O_RDONLY, 0666); err != nil {
		return "", NewFileReadError(file.Name())
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
		reg := utils.RegexpParser(g.Config.DataPartten)
		list := reg.FindAllStringSubmatch(string(line), -1)
		if len(list) != 0 {
			match = list[0][0]
			break
		}
	}

	if match == "" {
		return "", NewDataFilePathParseError(file.Name())
	}
	cacheDir := filepath.Join(match, "webCaches")
	dataPath := filepath.Join(cacheDir, g.GetVersionDir(cacheDir), baseData)

	var fi os.FileInfo
	if fi, err = os.Stat(dataPath); err != nil {
		return "", NewFileNotExistError(fi.Name())
	}
	expTime := int64(2 * 3600)
	modTime := fi.ModTime().Unix()
	now := time.Now().Unix()
	if now-modTime > expTime {
		return "", NewDataFileExpiredError(fi.Name(), g.GachaType)
	}
	return utils.CopyFileToTemp(dataPath), nil
}

func (g *Gacha) ParseGachaLink(dataPath string) (string, error) {
	var data []byte
	var err error
	if data, err = os.ReadFile(dataPath); err != nil {
		return "", NewFileReadError(dataPath)
	}
	reg := utils.RegexpParser(g.Config.LinkPartten)
	match := reg.FindAllStringSubmatch(string(data), -1)
	if len(match) == 0 {
		return "", NewLinkParseError(g.GachaType)
	}
	return match[len(match)-1][0], nil
}

func (g *Gacha) GetGachaLink() error {
	var err error
	var logPath string
	var dataPath string
	if logPath, err = g.GetLogPath(g.Config.BaseLogPath); err != nil {
		return err
	}
	if dataPath, err = g.GetDataPath(logPath, g.Config.BaseDataPath); err != nil {
		return err
	}
	// 返回链接
	if g.Link, err = g.ParseGachaLink(dataPath); err != nil {
		return err
	}
	g.CleanTempFile(logPath, dataPath)
	return nil
}

func (g *Gacha) CleanTempFile(files ...string) {
	utils.CleanTempFile(files...)
}

func (g *Gacha) GetVersionDir(cacheDir string) string {
	var cachePathArr []string
	var dir *os.File
	var err error
	var ens []os.DirEntry
	if dir, err = os.Open(cacheDir); err != nil {
		log.Println(err)
		return ""
	}
	defer func(dir *os.File) {
		if err = dir.Close(); err != nil {
			log.Println(err)
			return
		}
	}(dir)
	if ens, err = dir.ReadDir(0); err != nil {
		log.Println(err)
		return ""
	}

	for _, entry := range ens {
		if entry.IsDir() {
			reg := utils.RegexpParser("^\\d+\\.\\d+\\.\\d+\\.\\d+")
			if reg.MatchString(entry.Name()) {
				cachePathArr = append(cachePathArr, entry.Name())
			}
		}
	}
	return utils.FindMaxVersion(cachePathArr)
}
