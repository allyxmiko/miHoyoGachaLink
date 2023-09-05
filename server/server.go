package server

import (
	"github.com/gofiber/fiber/v2"
	"miHoyoGachaLink/gacha"
)

type Resp struct {
	code int
	msg  string
	data interface{}
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
		return c.JSONP(r)
	})
	if err := app.Listen("127.0.0.1:64127"); err != nil {
		return
	}
}
