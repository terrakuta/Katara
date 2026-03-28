package anime_repo

import (
	"Katara/internal/adapters/documents"
	"Katara/internal/domain/anime"
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type AnimeRepo struct {
	db *mongo.Collection
}

func NewAnimeRepository(db *mongo.Database) *AnimeRepo {
	return &AnimeRepo{
		db: db.Collection("anime"),
	}
}

func (s *AnimeRepo) GetAnimeByID(ctx context.Context, AniListID int) (*anime.Anime, error) {
	var animeDocument documents.AnimeDocument

	err := s.db.FindOne(ctx, bson.M{"anilist_id": AniListID}).Decode(&animeDocument)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, anime.ErrNotFound
		}
		return nil, err
	}
	return animeDocument.ToDomain(), nil
}

func (s *AnimeRepo) SaveAnime(ctx context.Context, anime *anime.Anime) (*anime.Anime, error) {
	var animeDoc = documents.FromDomain(anime)
	_, err := s.db.InsertOne(ctx, animeDoc)

	if err != nil {
		return nil, err
	}
	savedAnime := animeDoc.ToDomain()

	return savedAnime, nil
}

func (s *AnimeRepo) UpsertAnime(ctx context.Context, a *anime.Anime) error {
	doc := documents.FromDomain(a)
	opts := options.UpdateOne().SetUpsert(true)
	_, err := s.db.UpdateOne(
		ctx,
		bson.M{"anilist_id": doc.AniListID},
		bson.M{"$setOnInsert": doc},
		opts,
	)
	return err
}

func (s *AnimeRepo) GetAnimeWithFilters(ctx context.Context, animeFilter anime.AnimeFilter) ([]anime.Anime, error) {
	var sort bson.D

	switch animeFilter.Sort {
	case anime.SortTrendingDesc:
		sort = bson.D{{"anime_trending", -1}}
	case anime.SortPopularityDesc:
		sort = bson.D{{"anime_popularity", -1}}
	case anime.SortScoreDesc:
		sort = bson.D{{"anime_average_score", -1}}
	}

	opts := options.Find()
	if sort != nil {
		opts.SetSort(sort)
	}

	filter := bson.M{}

	if animeFilter.Status != "" {
		filter["anime_status"] = animeFilter.Status
	}
	if animeFilter.Format != "" {
		filter["anime_format"] = animeFilter.Format
	}
	if animeFilter.Season != "" {
		filter["anime_season"] = animeFilter.Season
	}
	if animeFilter.Year != 0 {
		filter["anime_season_year"] = animeFilter.Year
	}
	if len(animeFilter.Genre) > 0 {
		filter["anime_genres"] = bson.M{"$in": animeFilter.Genre}
	}

	if animeFilter.Page < 1 {
		animeFilter.Page = 1
	}

	if animeFilter.PerPage < 1 {
		animeFilter.PerPage = 20
	}

	opts.SetLimit(int64(animeFilter.PerPage))
	opts.SetSkip(int64((animeFilter.Page - 1) * animeFilter.PerPage))

	cursor, err := s.db.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var docs []documents.AnimeDocument
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}

	result := make([]anime.Anime, len(docs))

	for i, doc := range docs {
		result[i] = *doc.ToDomain()
	}

	return result, nil
}

func (s *AnimeRepo) SaveMany(ctx context.Context, animes []anime.Anime) error {
	docs := make([]interface{}, len(animes))
	for i, a := range animes {
		docs[i] = documents.FromDomain(&a)
	}

	_, err := s.db.InsertMany(ctx, docs)
	return err
}
