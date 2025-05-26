package main

import (
	"net/http"

	handler "github.com/Auxesia23/url_shortener/internal/handlers"
	middlewares "github.com/Auxesia23/url_shortener/internal/middlewares"
	"github.com/gin-gonic/gin"
)


type application struct {
    config config
    userHandler handler.UserHandler
    urlHandler handler.UrlHandler
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
	r.GET("/:id")
	{
		v1 := r.Group("/v1")
		v1.GET("/status", func(c *gin.Context){
			c.JSON(200, gin.H{"status" : "OK, Server up and running"})
		})
		
		{
			auth := v1.Group("/auth")
			auth.GET("/google", app.userHandler.HandleGoogleLogin)
			auth.GET("/google/callback", app.userHandler.HandleGoogleCallback)
		}
		
		{
			urls := v1.Group("/urls")
			urls.Use(middlewares.JwtAuthMiddleware())
			urls.POST("/", app.urlHandler.HandleCreateUrl)
			urls.GET("/",  app.urlHandler.HandleGetUrlByEmail)
			urls.GET("/:id", app.urlHandler.HandleGetUrl)
			urls.DELETE("/:id", app.urlHandler.HandleDeleteUrl)
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
