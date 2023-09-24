package server

import (
	"github.com/gofiber/fiber/v2"
	"miHoyoGachaLink/gacha"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Start() {
	app := fiber.New()
	app.Get("/gacha/:gtype", func(c *fiber.Ctx) error {
		gtype := c.Params("gtype")
		gachc := gacha.NewGacha(gtype)
		r := &Resp{
			200,
			"请求成功！",
			gachc.GetGachaLink(),
		}
		return c.JSON(r)
	})
	if err := app.Listen("127.0.0.1:64127"); err != nil {
		return
	}
}
