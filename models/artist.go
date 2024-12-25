package models

type Artist struct {
	ArtistId string   `json:"artist_id"`
	MasterId int      `json:"master_id"`
	Name     string   `json:"name"`
	Genre    string   `json:"genre"`
	Members  []string `json:"members"`
	Websites []string `json:"websites"`
	Albums   []Album  `json:"albums"`
}
