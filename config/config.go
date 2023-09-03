package config

import "miHoyoGachaLink/constant"

type GachaConfig struct {
	BaseLogPath  string
	BaseDataPath string
	DataPartten  string
	LinkPartten  string
}

func NewGachaConfig(configType string) *GachaConfig {
	switch configType {
	case constant.Genshin:
		return newGenshinConfig()
	case constant.StarRail:
		return newStarRailConfig()
	}
	panic("不支持的配置类型")
}
