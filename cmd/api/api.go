package main

import (
	"net/http"

	handler "github.com/Auxesia23/url_shortener/internal/handlers"
	"github.com/gin-gonic/gin"
)


type application struct {
    config config
    userHandler handler.UserHandler
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
	{
		v1 := r.Group("/v1")
		v1.GET("/status", func(c *gin.Context){
			c.JSON(200, gin.H{"status" : "OK, Server up and running"})
		})
		
		{
			auth := v1.Group("/auth")
			auth.GET("/google", app.userHandler.GoogleLogin)
			auth.GET("/google/callback", app.userHandler.GoogleCallback)
		}
	}
	

	return r
}

func (app *application) run(mux http.Handler) error {
   server := &http.Server{
       Addr:    app.config.addr,
       Handler: mux,
   }

   return server.ListenAndServe()
}
