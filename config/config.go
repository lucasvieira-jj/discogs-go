package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	MaxArtists = 10
	MaxAlbums  = 10
	MaxTracks  = 10
	UserAgent  = "LucasVieira/1.0"
	TOKEN      string
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
