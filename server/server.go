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
	var err error
	var gachc *gacha.Gacha
	var jsonb *Resp
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		GETOnly:               true,
	})
	app.Get("/gacha/:gtype", func(c *fiber.Ctx) error {
		gtype := c.Params("gtype")
		if gachc, err = gacha.NewGacha(gtype); err != nil {
			jsonb = &Resp{
				500,
				err.Error(),
				nil,
			}
		} else {
			jsonb = &Resp{
				0,
				"请求成功！",
				gachc.Link,
			}
		}
		return c.JSON(jsonb)
	})
	if err = app.Listen("127.0.0.1:64127"); err != nil {
		return
	}
}
