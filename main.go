package main

import (
	"fmt"
	"miHoyoGachaLink/constant"
	"miHoyoGachaLink/gacha"
)

func main() {
	sr := gacha.NewGacha(constant.StarRail)
	fmt.Println(sr.GachaLink)
}
