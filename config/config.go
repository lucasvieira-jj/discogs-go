package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

var (
	AppName = "Discogs App Crawler"

	BaseUrl           = "https://www.discogs.com"
	MaxArtists        = 10
	MaxAlbums         = 2
	UserAgent         = "LucasVieira/1.0"
	TOKEN             string
	GenreToSearch     = "hip hop"
	MaxTextSize       = 50
	MaxRetries        = 10
	WaitTimeExecution = 5 * time.Second
)

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	TOKEN = os.Getenv("TOKEN")
	if TOKEN == "" {
		log.Fatal("You need to set env TOKEN to run this application")
	}

}
