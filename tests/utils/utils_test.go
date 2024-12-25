package utils

import (
	"github.com/lucasvieira-jj/discogs-go/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePrimaryKey(t *testing.T) {
	key1 := "123456789"
	key2 := "test"
	key3 := "2024"

	expected := "651913d93ad82fd9e812fd0b78fe9a3f"

	result := utils.CreatePrimaryKey(key1, key2, key3)

	assert.Equal(t, expected, result, "The generated primary key should be "+expected)
}

func TestGenerateFileName(t *testing.T) {
	key1 := "Genre Test"

	expected := "2024-12-25T18:44:25-discogsappcrawler-genretest"

	result := utils.GenerateFileName(key1)

	assert.Equal(t, expected, result, "The generated filename should be "+expected)
}
