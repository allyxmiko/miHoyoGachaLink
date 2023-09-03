package config

import (
	"path/filepath"
)

func newGenshinConfig() *GachaConfig {
	return &GachaConfig{
		BaseLogPath:  filepath.Join("AppData", "LocalLow", "miHoYo", "原神", "output_log.txt.last"),
		BaseDataPath: filepath.Join("webCaches", "2.15.0.0", "Cache", "Cache_Data", "data_2"),
		DataPartten:  "[A-Z]{1}:.*YuanShen_Data",
		LinkPartten:  "https://hk4e-api.mihoyo.com/event/gacha_info/api/getGachaLog[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]game_biz=hk4e_cn",
	}
}
