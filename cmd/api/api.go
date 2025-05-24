package main

import (
	"net/http"

	"github.com/Auxesia23/url_shortener/internal/handlers"
	"github.com/gin-gonic/gin"
)


type application struct {
    config config
    HealthCheck handler.HealthCheck
    User handler.UserHandler
}

type config struct {
    addr string
}

func NewApplication(cfg config) *application {
    return &application{
        config: cfg,
    }
}

func (app *application) mount() http.Handler {
	r := gin.Default()
	
	r.GET("/healthcheck", app.HealthCheck.Check)

	return r
}

func (app *application) run(mux http.Handler) error {
   server := &http.Server{
       Addr:    app.config.addr,
       Handler: mux,
   }

   return server.ListenAndServe()
}
