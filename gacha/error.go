package gacha

import "miHoyoGachaLink/constant"

type FileNotExistError struct {
	fileName string
}

func (f *FileNotExistError) Error() string {
	return f.fileName + " -----> 该文件不存在！"
}

func NewFileNotExistError(fileName string) error {
	return &FileNotExistError{fileName: fileName}
}

type FileReadError struct {
	fileName string
}

func (f *FileReadError) Error() string {
	return f.fileName + " -----> 该文件读取失败！"
}

func NewFileReadError(fileName string) error {
	return &FileReadError{fileName: fileName}
}

type DataFilePathParseError struct {
	fileName string
}

func (d *DataFilePathParseError) Error() string {
	return d.fileName + " -----> 数据文件路径解析失败！"
}

func NewDataFilePathParseError(fileName string) error {
	return &DataFilePathParseError{fileName: fileName}
}

type DataFileExpiredError struct {
	fileName string
	gtype    string
}

func (d *DataFileExpiredError) Error() string {
	var plt string
	switch d.gtype {
	case constant.Genshin:
		plt = "原神"
	case constant.StarRail:
		plt = "星穹铁道"
	default:
		plt = "游戏"
	}
	return d.fileName + " -----> 数据文件已过期！请登录" + plt + "查看抽卡记录以刷新有效期！"
}

func NewDataFileExpiredError(fileName, gtype string) error {
	return &DataFileExpiredError{fileName: fileName, gtype: gtype}
}

type LinkParseError struct {
	gtype string
}

func (l *LinkParseError) Error() string {
	return "未匹配到抽卡链接!请登录" + l.gtype + "查看抽卡记录以刷新链接！"
}

func NewLinkParseError(gtype string) error {
	return &LinkParseError{gtype: gtype}
}
