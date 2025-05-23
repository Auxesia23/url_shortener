package main

import (
	"log"

	"github.com/Auxesia23/url_shortener/internal/db"
	"github.com/Auxesia23/url_shortener/internal/handler"
	"github.com/Auxesia23/url_shortener/internal/repository"
	"github.com/Auxesia23/url_shortener/internal/service"
)

func main() {
	//Server configuration
	cfg := config{
		addr: ":8080",
	}
	
	
	//Database
	db, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	
	//Repository
	userRepo := repository.NewUserRepository(db)
	
	//service
	userService := service.NewUserService(userRepo)
	
	//Handler
	healthCheckHandler := handler.NewHealthCheck()
	userHandler := handler.NewUserHandler(*userService)
	
	
	//Dependencies Injection for Application
	app := &application{
		config: cfg,
		HealthCheck: *healthCheckHandler,
		User: *userHandler,
	}
	
	//Initiate Handlers
	mux := app.mount()
	
	//Run server with provided handlers
	log.Fatal(app.run(mux))
}
