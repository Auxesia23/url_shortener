package main

import (
	"net/http"
	"time"

	handler "github.com/Auxesia23/url_shortener/internal/handlers"
	middlewares "github.com/Auxesia23/url_shortener/internal/middlewares"
	"github.com/gin-gonic/gin"
)


type application struct {
    config config
    userHandler handler.UserHandler
    urlHandler handler.UrlHandler
    redirectHandler handler.RedirectHandler
}


type config struct {
    addr string
    readTimeout time.Duration
    writeTimeout time.Duration
    idleTimeout time.Duration
}

func NewApplication(cfg config) *application {
    return &application{
        config: cfg,
    }
}

func (app *application) mount() http.Handler {
	r := gin.Default()
	
	r.GET("/:id",app.redirectHandler.HandleRedirect)
	
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
       ReadTimeout: app.config.readTimeout,
       WriteTimeout: app.config.writeTimeout,
       IdleTimeout: app.config.idleTimeout,
   }

   return server.ListenAndServe()
}
