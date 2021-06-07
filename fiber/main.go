package main

import (
	"log"
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	app := fiber.New()

	app.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello!")
	})

	nrapp, err := newrelic.NewApplication()
	if err != nil {
		log.Println(err)
	}

	http.HandleFunc(newrelic.WrapHandleFunc(nrapp, "/", adaptor.FiberApp(app)))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
