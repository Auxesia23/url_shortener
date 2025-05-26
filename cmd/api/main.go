package main

import (
	"log"

	"github.com/Auxesia23/url_shortener/internal/auth"
	"github.com/Auxesia23/url_shortener/internal/db"
	handler "github.com/Auxesia23/url_shortener/internal/handlers"
	"github.com/Auxesia23/url_shortener/internal/repositories"
	"github.com/Auxesia23/url_shortener/internal/services"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file")
	}
	
	//Server configuration
	cfg := config{
		addr: ":8080",
	}
	
	//Initialize Oauth
	auth.InitOauth()
	
	//Database
	db, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	
	//Repository
	userRepo := repository.NewUserRepository(db)
	urlRepository := repository.NewUrlRepository(db)
	
	//service
	userService := service.NewUserService(userRepo)
	urlService := service.NewUrlService(urlRepository)
	
	//Handler
	userHandler := handler.NewUserHandler(userService)
	urlHandler := handler.NewUrlHandler(urlService)
	
	//Dependencies Injection for Application
	app := &application{
		config: cfg,
		userHandler: userHandler,
		urlHandler: urlHandler,
	}
	
	//Initiate Handlers
	mux := app.mount()
	
	//Run server with provided handlers
	log.Fatal(app.run(mux))
}
