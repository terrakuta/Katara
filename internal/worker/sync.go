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
	total := 0

	err := w.provider.FetchAll(ctx, func(batch []anime.Anime) error {
		for _, a := range batch {
			if err := w.repo.UpsertAnime(ctx, &a); err != nil {
				log.Printf("upsert error for anime %d: %v", a.AniListID, err)
			}
		}
		total += len(batch)
		log.Printf("batch saved, total so far: %d", total)
		return nil
	})

	if err != nil {
		return err
	}

	log.Printf("sync complete: %d anime", total)
	total = 0
	return nil
}
