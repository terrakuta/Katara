package anime

import (
	"context"
	"errors"
	"log"
)

var ErrNotFound = errors.New("not found")

type AnimeService struct {
	repo AnimeRepository
}

func NewAnimeService(repo AnimeRepository) *AnimeService {
	log.Printf("NewAnimeService called, repo = %v", repo)
	return &AnimeService{repo: repo}
}

func (s *AnimeService) GetAnimeByID(ctx context.Context, AniListID int) (*Anime, error) {

	result, err := s.repo.GetAnimeByID(ctx, AniListID)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *AnimeService) GetAnimeWithFilters(ctx context.Context, animeFilter AnimeFilter) ([]Anime, error) {
	return s.repo.GetAnimeWithFilters(ctx, animeFilter)
}
