package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()

    app.Static("/static", "static/content")

    app.Get("/:neededVal/:optVal?/*", func(c *fiber.Ctx) error {
        return c.SendString("neededVal: " + c.Params("neededVal") + ", optVal: " + c.Params("optVal") + ", wildcard: " + c.Params("*"))
    })

    app.Listen(":3000")
}