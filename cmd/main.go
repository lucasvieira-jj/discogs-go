package main

import (
	"fmt"
	"github.com/lucasvieira-jj/discogs-go/config"
	"github.com/lucasvieira-jj/discogs-go/internal/pipeline"
	"github.com/lucasvieira-jj/discogs-go/internal/scraper"
	"github.com/lucasvieira-jj/discogs-go/internal/utils"
	"log"
)

func main() {
	config.InitConfig()

	log.Printf("Discogs is starting...")
	client := scraper.NewClient()
	artists := client.ArtistsRetrieved(config.GenreToSearch)

	for i, artist := range artists {
		artists[i] = client.ArtistsSearchAlbums(artist)
		artists[i] = client.ArtistsSearchTracks(artists[i])
	}

	fmt.Println("Artists:", utils.JsonConverter(artists))

	pipeline.SaveRawData(artists)
	pipeline.SaveToJSONL(artists)
}
