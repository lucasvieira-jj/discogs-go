package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/lucasvieira-jj/discogs-go/config"
	"github.com/lucasvieira-jj/discogs-go/models"
	"strings"
	"time"
)

func JsonConverter(list []models.Artist) string {
	convertedArtists, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return ""
	}

	return string(convertedArtists)
}

func CreatePrimaryKey(key1, key2, key3 string) string {
	data := config.AppName + "&" + key1 + key2 + key3
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GenerateFileName(key1 string) string {
	appname := strings.ReplaceAll(strings.ToLower(config.AppName), " ", "")
	key1 = strings.ReplaceAll(strings.ToLower(key1), " ", "")

	now := time.Now()

	dateFormated := now.Format("2006-01-02T15:04:05")
	fileName := dateFormated + "-" + appname + "-" + key1 + ".json"
	return fileName
}
