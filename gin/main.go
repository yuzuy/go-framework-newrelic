package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	r := gin.New()
	app, err := newrelic.NewApplication()
	if err != nil {
		log.Println(err)
	}
	r.Use(nrgin.Middleware(app))
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
