package worker

import (
	"Katara/internal/domain/anime"
	"context"
	"log"
	"time"
)

type SyncWorker struct {
	provider anime.AnimeProvider
	repo     anime.AnimeRepository
}

func NewSyncWorker(provider anime.AnimeProvider, repo anime.AnimeRepository) *SyncWorker {
	return &SyncWorker{
		provider: provider,
		repo:     repo,
	}
}

func (w *SyncWorker) Start() {
	go func() {
		for {
			if err := w.sync(); err != nil {
				log.Printf("sync error: %v", err)
			}
			time.Sleep(24 * time.Hour)
		}
	}()
}

func (w *SyncWorker) sync() error {
	ctx := context.Background()

	animes, err := w.provider.FetchAll(ctx)
	if err != nil {
		return err
	}

	for _, a := range animes {
		if err := w.repo.UpsertAnime(ctx, &a); err != nil {
			log.Printf("upsert error for anime %d: %v", a.AniListID, err)
		}
	}

	log.Printf("sync complete: %d anime updated", len(animes))
	return nil
}
