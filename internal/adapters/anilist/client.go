package anilist

import (
	"Katara/internal/domain/anime"
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type AnilistClient struct {
	httpClient *http.Client
	baseURL    string
}

func NewAnilistClient() *AnilistClient {
	return &AnilistClient{
		httpClient: &http.Client{Timeout: 30 * time.Second},
		baseURL:    "https://graphql.anilist.co",
	}
}

func (c *AnilistClient) FetchAll(ctx context.Context, onBatch func([]anime.Anime) error) error {
	query := `
    query ($page: Int) {
        Page(page: $page, perPage: 50) {
            media(type: ANIME) {
                id
                title { romaji english native userPreferred }
                episodes
                genres
                averageScore
                popularity
                trending
                status
                format
                season
                seasonYear
                description
                coverImage { extraLarge large medium color }
                trailer { id site thumbnail }
                studios { nodes { id name } }
            }
            pageInfo { hasNextPage }
        }
    }`

	page := 1

	for {
		log.Printf("fetching page %d", page)

		body := map[string]interface{}{
			"query":     query,
			"variables": map[string]interface{}{"page": page},
		}

		jsonBody, err := json.Marshal(body)
		if err != nil {
			return err
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL, bytes.NewBuffer(jsonBody))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return err
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			resp.Body.Close()
			log.Printf("rate limited, waiting 60s...")
			time.Sleep(60 * time.Second)
			continue
		}

		var result AnilistResponse
		err = json.NewDecoder(resp.Body).Decode(&result)
		resp.Body.Close()
		if err != nil {
			return err
		}

		log.Printf("page %d: got %d anime, hasNextPage: %v",
			page,
			len(result.Data.Page.Media),
			result.Data.Page.PageInfo.HasNextPage,
		)

		if len(result.Errors) > 0 {
			log.Printf("anilist errors on page %d: %+v", page, result.Errors)
			time.Sleep(5 * time.Second)
			continue
		}

		if len(result.Data.Page.Media) == 0 && page > 1 {
			log.Printf("page %d returned 0 anime, retrying...", page)
			time.Sleep(5 * time.Second)
			continue
		}

		if err := onBatch(result.toDomain()); err != nil {
			return err
		}

		if !result.Data.Page.PageInfo.HasNextPage {
			break
		}
		page++
		time.Sleep(700 * time.Millisecond)
	}

	return nil
}

type AnilistResponse struct {
	Errors []struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	} `json:"errors"`
	Data struct {
		Page struct {
			Media []struct {
				ID    int `json:"id"`
				Title struct {
					Romaji        string `json:"romaji"`
					English       string `json:"english"`
					Native        string `json:"native"`
					UserPreferred string `json:"userPreferred"`
				} `json:"title"`
				Episodes     int      `json:"episodes"`
				Genres       []string `json:"genres"`
				AverageScore int      `json:"averageScore"`
				Popularity   int      `json:"popularity"`
				Trending     int      `json:"trending"`
				Status       string   `json:"status"`
				Format       string   `json:"format"`
				Season       string   `json:"season"`
				SeasonYear   int      `json:"seasonYear"`
				Description  string   `json:"description"`
				CoverImage   struct {
					ExtraLarge string `json:"extraLarge"`
					Large      string `json:"large"`
					Medium     string `json:"medium"`
					Color      string `json:"color"`
				} `json:"coverImage"`
				Trailer struct {
					ID        string `json:"id"`
					Site      string `json:"site"`
					Thumbnail string `json:"thumbnail"`
				} `json:"trailer"`
				Tags []struct {
					ID               int    `json:"id"`
					Name             string `json:"name"`
					Description      string `json:"description"`
					Category         string `json:"category"`
					Rank             int    `json:"rank"`
					IsGeneralSpoiler bool   `json:"isGeneralSpoiler"`
					IsMediaSpoiler   bool   `json:"isMediaSpoiler"`
					IsAdult          bool   `json:"isAdult"`
					UserID           int    `json:"userId"`
				} `json:"tags"`
				Studios struct {
					Nodes []struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
					} `json:"nodes"`
				} `json:"studios"`
			} `json:"media"`
			PageInfo struct {
				HasNextPage bool `json:"hasNextPage"`
			} `json:"pageInfo"`
		} `json:"Page"`
	} `json:"data"`
}

func (r *AnilistResponse) toDomain() []anime.Anime {
	result := make([]anime.Anime, len(r.Data.Page.Media))
	for i, m := range r.Data.Page.Media {
		studios := make([]anime.Studio, len(m.Studios.Nodes))
		for j, s := range m.Studios.Nodes {
			studios[j] = anime.Studio{
				ID:   strconv.Itoa(s.ID),
				Name: s.Name,
			}
		}

		tags := make([]anime.Mediatag, len(m.Tags))
		for j, t := range m.Tags {
			tags[j] = anime.Mediatag{
				ID:               t.ID,
				Name:             t.Name,
				Description:      t.Description,
				Category:         t.Category,
				Rank:             t.Rank,
				IsGeneralSpoiler: t.IsGeneralSpoiler,
				IsMediaSpoiler:   t.IsMediaSpoiler,
				IsAdult:          t.IsAdult,
				UserID:           t.UserID,
			}
		}

		result[i] = anime.Anime{
			AniListID: m.ID,
			Title: anime.MediaTitle{
				Romaji:        m.Title.Romaji,
				English:       m.Title.English,
				Native:        m.Title.Native,
				UserPreferred: m.Title.UserPreferred,
			},
			CoverImage: anime.MediaCoverImage{
				ExtraLarge: m.CoverImage.ExtraLarge,
				Large:      m.CoverImage.Large,
				Medium:     m.CoverImage.Medium,
				Color:      m.CoverImage.Color,
			},
			Trailer: anime.MediaTrailer{
				ID:        m.Trailer.ID,
				Site:      m.Trailer.Site,
				Thumbnail: m.Trailer.Thumbnail,
			},
			Format:       anime.MediaFormat(m.Format),
			Status:       anime.MediaStatus(m.Status),
			Season:       anime.MediaSeason(m.Season),
			Episodes:     m.Episodes,
			Genres:       m.Genres,
			AverageScore: m.AverageScore,
			Popularity:   m.Popularity,
			Description:  m.Description,
			Trending:     m.Trending,
			SeasonYear:   m.SeasonYear,
			Studios:      studios,
			Tags:         tags,
		}
	}
	return result
}
