package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"anime-discovery-api/config"
	"anime-discovery-api/models"
)

var cfg *config.Config
var cache = make(map[string]cacheEntry)

type cacheEntry struct {
	data       []byte
	expiration time.Time
}

func Init(c *config.Config) {
	cfg = c
}

func FetchTopAnime(page int) ([]models.Anime, error) {
	cacheKey := fmt.Sprintf("top_anime_%d", page)

	if cached, ok := cache[cacheKey]; ok && time.Now().Before(cached.expiration) {
		var animeList []models.Anime
		err := json.Unmarshal(cached.data, &animeList)
		if err == nil {
			return animeList, nil
		}
	}

	url := fmt.Sprintf("%s/top/anime?page=%d", cfg.JikanAPIURL, page)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var jikanResponse struct {
		Data []struct {
			MalID        int    `json:"mal_id"`
			Title        string `json:"title"`
			TitleEnglish string `json:"title_english"`
			Synopsis     string `json:"synopsis"`
			Images       struct {
				JPG struct {
					ImageURL string `json:"image_url"`
				} `json:"jpg"`
			} `json:"images"`
			Score    float64 `json:"score"`
			Episodes int     `json:"episodes"`
			Rating   string  `json:"rating"`
			Year     int     `json:"year"`
			Status   string  `json:"status"`
			Genres   []struct {
				Name string `json:"name"`
			} `json:"genres"`
		} `json:"data"`
	}

	err = json.Unmarshal(body, &jikanResponse)
	if err != nil {
		return nil, err
	}

	animeList := make([]models.Anime, 0, len(jikanResponse.Data))
	for _, a := range jikanResponse.Data {
		genre := ""
		if len(a.Genres) > 0 {
			genre = a.Genres[0].Name
		}

		titleEnglish := a.TitleEnglish
		if titleEnglish == "" {
			titleEnglish = a.Title
		}

		animeList = append(animeList, models.Anime{
			ID:           a.MalID,
			Title:        a.Title,
			TitleEnglish: titleEnglish,
			Synopsis:     a.Synopsis,
			ImageURL:     a.Images.JPG.ImageURL,
			Score:        a.Score,
			Episodes:     a.Episodes,
			Rating:       a.Rating,
			Year:         a.Year,
			Status:       a.Status,
			Genre:        genre,
		})
	}

	cacheDuration := time.Duration(cfg.CacheTimeout) * time.Minute
	cache[cacheKey] = cacheEntry{
		data:       body,
		expiration: time.Now().Add(cacheDuration),
	}

	return animeList, nil
}

func FetchAnimeByID(id int) (*models.Anime, error) {
	cacheKey := fmt.Sprintf("anime_%d", id)

	if cached, ok := cache[cacheKey]; ok && time.Now().Before(cached.expiration) {
		var anime models.Anime
		err := json.Unmarshal(cached.data, &anime)
		if err == nil {
			return &anime, nil
		}
	}

	url := fmt.Sprintf("%s/anime/%d", cfg.JikanAPIURL, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var jikanResponse struct {
		Data struct {
			MalID        int    `json:"mal_id"`
			Title        string `json:"title"`
			TitleEnglish string `json:"title_english"`
			Synopsis     string `json:"synopsis"`
			Images       struct {
				JPG struct {
					ImageURL string `json:"image_url"`
				} `json:"jpg"`
			} `json:"images"`
			Score    float64 `json:"score"`
			Episodes int     `json:"episodes"`
			Rating   string  `json:"rating"`
			Year     int     `json:"year"`
			Status   string  `json:"status"`
			Genres   []struct {
				Name string `json:"name"`
			} `json:"genres"`
		} `json:"data"`
	}

	err = json.Unmarshal(body, &jikanResponse)
	if err != nil {
		return nil, err
	}

	genre := ""
	if len(jikanResponse.Data.Genres) > 0 {
		genre = jikanResponse.Data.Genres[0].Name
	}

	titleEnglish := jikanResponse.Data.TitleEnglish
	if titleEnglish == "" {
		titleEnglish = jikanResponse.Data.Title
	}

	anime := &models.Anime{
		ID:           jikanResponse.Data.MalID,
		Title:        jikanResponse.Data.Title,
		TitleEnglish: titleEnglish,
		Synopsis:     jikanResponse.Data.Synopsis,
		ImageURL:     jikanResponse.Data.Images.JPG.ImageURL,
		Score:        jikanResponse.Data.Score,
		Episodes:     jikanResponse.Data.Episodes,
		Rating:       jikanResponse.Data.Rating,
		Year:         jikanResponse.Data.Year,
		Status:       jikanResponse.Data.Status,
		Genre:        genre,
	}

	cacheDuration := time.Duration(cfg.CacheTimeout) * time.Minute
	cache[cacheKey] = cacheEntry{
		data:       body,
		expiration: time.Now().Add(cacheDuration),
	}

	return anime, nil
}
