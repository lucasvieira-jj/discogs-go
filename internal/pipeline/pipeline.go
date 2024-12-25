package pipeline

import (
	"encoding/json"
	"fmt"
	"github.com/lucasvieira-jj/discogs-go/config"
	"github.com/lucasvieira-jj/discogs-go/internal/utils"
	"github.com/lucasvieira-jj/discogs-go/models"
	"log"
	"os"
)

func SaveRawData(data []models.Artist) {
	log.Println("Saving the data into raw")

	dirPath := "./internal/storage/raw/"
	fileName := utils.GenerateFileName(config.GenreToSearch)
	filePath := dirPath + fileName

	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Println("Error when try to create a file:", err)
		return
	}

	fmt.Println("Saving raw data to: " + filePath)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("There is a problem to create a file", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("There is a problem to close a file", err)
		}
	}(file)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("There is a problem to save file json:", err)
		return
	}

	fmt.Println("The data is saved with success")

}

func SaveToJSONL(data []models.Artist) {
	log.Println("Saving the formatted data into the trusted layer")

	dirPath := "./internal/storage/trusted/"
	fileName := utils.GenerateFileName(config.GenreToSearch)
	filePath := dirPath + fileName

	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Println("Error when try to create a file:", err)
		return
	}

	fmt.Println("Saving data to: " + filePath)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("There is a problem to create a file", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("There is a problem to close a file", err)
		}
	}(file)

	for _, item := range data {
		line, err := json.Marshal(item)
		if err != nil {
			fmt.Println("Error to save the data into file", err)
			return
		}

		_, err = file.Write(append(line, '\n'))
		if err != nil {
			fmt.Println("Error to write data into file:", err)
			return
		}
	}

	fmt.Println("The data is saved with success")

}
