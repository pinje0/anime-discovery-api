package models

type Anime struct {
	ID           int     `json:"id"`
	Title        string  `json:"title"`
	TitleEnglish string  `json:"title_english"`
	Synopsis     string  `json:"synopsis"`
	ImageURL     string  `json:"image_url"`
	Score        float64 `json:"score"`
	Episodes     int     `json:"episodes"`
	Rating       string  `json:"rating"`
	Year         int     `json:"year"`
	Status       string  `json:"status"`
	Genre        string  `json:"genre"`
}

type AnimeListResponse struct {
	Anime []Anime `json:"anime"`
	Page  int     `json:"page"`
	Total int     `json:"total"`
}
