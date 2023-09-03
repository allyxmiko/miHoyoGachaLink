package config

import "path/filepath"

func newStarRailConfig() *GachaConfig {
	return &GachaConfig{
		BaseLogPath:  filepath.Join("AppData", "LocalLow", "miHoYo", "崩坏：星穹铁道", "Player.log"),
		BaseDataPath: filepath.Join("webCaches", "2.15.0.0", "Cache", "Cache_Data", "data_2"),
		DataPartten:  "[A-Z]{1}:.*starrail_Data",
		LinkPartten:  "https://api-takumi.mihoyo.com/common/gacha_record/api/getGachaLog[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]plat_type=pc",
	}
}
