package anime

import (
	"context"
)

type AnimeRepository interface {
	GetAnimeByID(ctx context.Context, AniListID int) (*Anime, error)
	GetAnimeWithFilters(ctx context.Context, animeFilter AnimeFilter) ([]Anime, error)
	SaveAnime(ctx context.Context, anime *Anime) (*Anime, error)
	SaveMany(ctx context.Context, animes []Anime) error
	UpsertAnime(ctx context.Context, anime *Anime) error
}

type AnimeProvider interface {
	FetchAll(ctx context.Context) ([]Anime, error)
}
