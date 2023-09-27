package main

import (
	"fmt"
	"miHoyoGachaLink/constant"
	"miHoyoGachaLink/server"
)

func main() {
	fmt.Println("当前版本：v" + constant.Version)
	fmt.Println("1. 点击链接前请先登录游戏查看一下对应游戏的抽卡记录。")
	fmt.Println("2. 查看记录后请确保游戏关闭。")
	fmt.Println("3. 点击以下对应链接获取对应的抽卡链接。")
	fmt.Println("原神: http://127.0.0.1:64127/gacha/ys")
	fmt.Println("星穹铁道: http://127.0.0.1:64127/gacha/sr")
	server.Start()
}
