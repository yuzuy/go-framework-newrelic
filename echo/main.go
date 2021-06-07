package main

import (
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	e := echo.New()
	app, err := newrelic.NewApplication()
	if err != nil {
		e.Logger.Warn(err)
	}
	e.Use(nrecho.Middleware(app))
	e.Logger.Fatal(e.Start(":8080"))
}
