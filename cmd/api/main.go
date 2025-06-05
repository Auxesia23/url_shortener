package main

import (
	"log"
	"os"
	"time"

	"github.com/Auxesia23/url_shortener/internal/auth"
	"github.com/Auxesia23/url_shortener/internal/db"
	handler "github.com/Auxesia23/url_shortener/internal/handlers"
	"github.com/Auxesia23/url_shortener/internal/repositories"
	"github.com/Auxesia23/url_shortener/internal/services"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file")
	}
	
	//Initialize ipinfo client
	client := ipinfo.NewClient(nil, nil, os.Getenv("IPINFO_TOKEN"))
	
	
	//Server configuration
	cfg := config{
		addr: "0.0.0.0:8080",
		readTimeout: 5 * time.Second,
		writeTimeout: 10 * time.Second,
		idleTimeout: 60 * time.Second,
	}
	
	//Initialize Oauth
	auth.InitOauth()
	
	//Database
	db, err := db.InitPostgres()
	if err != nil {
		log.Fatal(err)
	}
	
	//Repository
	userRepo := repository.NewUserRepository(db)
	urlRepository := repository.NewUrlRepository(db)
	analyticRepo := repository.NewAnalyticRepository(db)
	
	//service
	userService := service.NewUserService(userRepo)
	urlService := service.NewUrlService(urlRepository)
	redirectService := service.NewRedirectService(urlRepository)
	analyticService := service.NewAnalyticService(analyticRepo,client)
	
	//Handler
	userHandler := handler.NewUserHandler(userService)
	urlHandler := handler.NewUrlHandler(urlService,analyticService)
	redirectHandler := handler.NewRedirectHandler(redirectService,analyticService)
	
	//Dependencies Injection for Application
	app := &application{
		config: cfg,
		userHandler: userHandler,
		urlHandler: urlHandler,
		redirectHandler: redirectHandler,
	}
	
	//Initiate Handlers
	mux := app.mount()
	
	//Run server with provided handlers
	log.Fatal(app.run(mux))
}
