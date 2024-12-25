package models

type Album struct {
	MasterId int      `json:"master_id"`
	Name     string   `json:"name"`
	Year     string   `json:"year"`
	Record   []string `json:"Record"`
	Styles   []string `json:"styles"`
	Tracks   []Track  `json:"tracks"`
}
