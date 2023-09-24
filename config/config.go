package config

import "miHoyoGachaLink/constant"

type GachaConfig struct {
	BaseLogPath  string
	BaseDataPath string
	DataPartten  string
	LinkPartten  string
}

type SupportedTypeError struct {
}

func (s *SupportedTypeError) Error() string {
	return "不支持的配置类型"
}

func NewSupportedTypeError() error {
	return &SupportedTypeError{}
}

func NewGachaConfig(configType string) (*GachaConfig, error) {
	switch configType {
	case constant.Genshin:
		return newGenshinConfig(), nil
	case constant.StarRail:
		return newStarRailConfig(), nil
	default:
		return nil, NewSupportedTypeError()
	}
}
