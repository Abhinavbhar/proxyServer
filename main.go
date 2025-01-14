package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func main() {
	app := fiber.New()

	app.All("/", func(c *fiber.Ctx) error {
		fmt.Println(c.Request())
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)

		c.Request().CopyTo(req)
		res := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(res)
		if err := fasthttp.Do(req, res); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Proxy failed: " + err.Error())
		}
		c.Response().SetBodyRaw(res.Body())
		c.Response().SetStatusCode(res.StatusCode())
		return nil
	})

	app.Listen("0.0.0.0:8080")
}
