package scraper

import (
	"github.com/irlndts/go-discogs"
	"github.com/lucasvieira-jj/discogs-go/config"
	"github.com/lucasvieira-jj/discogs-go/internal/utils"
	"github.com/lucasvieira-jj/discogs-go/models"
	"log"
	"strconv"
	"strings"
	"time"
)

type ClientAPI struct {
	client discogs.Discogs
}

func NewClient() *ClientAPI {
	client, err := discogs.New(&discogs.Options{
		UserAgent: config.UserAgent,
		Currency:  "USD",
		Token:     config.TOKEN,
	})
	if err != nil {
		log.Fatalf("Error creating discogs client %v", err)
	}

	return &ClientAPI{client: client}
}

func (c *ClientAPI) ArtistsRetrieved(genre string) []models.Artist {
	var artists []models.Artist
	uniqueArtists := make(map[string]struct{})
	pagination := 1

	for len(artists) < config.MaxArtists {
		request := discogs.SearchRequest{
			Genre: genre,
			Page:  pagination,
			Type:  "release",
		}

		searchResult, _ := c.client.Search(request)

		for _, result := range searchResult.Results {
			if len(artists) >= config.MaxArtists {
				break
			}

			if _, exists := uniqueArtists[result.Title]; !exists {
				uniqueArtists[result.Title] = struct{}{}

				artistId := utils.CreatePrimaryKey(
					strconv.Itoa(result.MasterID), result.Title, result.Year,
				)
				artistName := strings.TrimSpace(strings.Split(result.Title, "-")[0])

				artist := models.Artist{
					ArtistId: artistId,
					MasterId: result.MasterID,
					Name:     artistName,
					Genre:    genre,
				}

				artists = append(artists, artist)
			}
		}

		pagination++
	}

	return artists
}

func (c *ClientAPI) ArtistsSearchAlbums(artist models.Artist) models.Artist {
	var albums []models.Album
	waitTime := config.WaitTimeExecution
	uniqueAlbums := make(map[string]struct{})

	request := discogs.SearchRequest{
		Artist: artist.Name,
		Type:   "release",
		Format: "album",
	}

	var err error
	var searchResult *discogs.Search

	for attempts := 0; attempts < config.MaxRetries; attempts++ {
		searchResult, err = c.client.Search(request)
		if err != nil {
			if attempts < config.MaxRetries-1 {
				log.Printf("Error searching artists %v. Retrying...\n", err)
				time.Sleep(waitTime)
				waitTime *= 2
			} else {
				log.Printf("Final attempt failed: %v\n", err)
			}
		} else {
			break
		}
	}

	if err == nil && searchResult != nil {
		for _, result := range searchResult.Results {
			if len(albums) >= config.MaxAlbums {
				break
			}

			if _, exists := uniqueAlbums[result.Title]; !exists {
				uniqueAlbums[result.Title] = struct{}{}

				limitedTitle := result.Title

				if len(result.Title) > config.MaxTextSize {
					limitedTitle = result.Title[:config.MaxTextSize]
				}

				album := models.Album{
					MasterId: result.MasterID,
					Name:     limitedTitle,
					Year:     result.Year,
					Styles:   result.Style,
					Record:   result.Label,
				}

				albums = append(albums, album)
			}
		}
	}

	artist.Albums = append(albums)
	return artist
}

func (c *ClientAPI) ArtistsSearchTracks(artist models.Artist) models.Artist {
	uniqueTracks := make(map[string]struct{})

	for i := range artist.Albums {
		album := &artist.Albums[i]

		searchResult, err := c.client.Master(artist.MasterId)
		if err != nil {
			log.Fatalf("Error searching artists %v", err)
		}

		var tracks []models.Track
		for trackPosition, result := range searchResult.Tracklist {
			if _, exists := uniqueTracks[result.Title]; !exists {
				uniqueTracks[result.Title] = struct{}{}

				track := models.Track{
					Number:   trackPosition + 1,
					Title:    result.Title,
					Duration: result.Duration,
				}

				tracks = append(tracks, track)
			}
		}

		album.Tracks = tracks

	}

	return artist
}
