package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type App struct {
	config *oauth2.Config
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET_ID")

	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:7000/auth/callback",
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	app := &App{config: conf}

	r := gin.Default()

	r.LoadHTMLFiles("index.html")

	r.GET("/auth/login", app.loginHandler)
	r.GET("/auth/oauth", app.oAuthHandler)
	r.GET("/auth/callback", app.oAuthCallbackHandler)

	if err := r.Run(":7000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
