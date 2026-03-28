package documents

import (
	"Katara/internal/domain/anime"
	"Katara/internal/domain/list"
	"Katara/internal/domain/user"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MediaTitle struct {
	Romaji        string `bson:"romaji"`
	English       string `bson:"english"`
	Native        string `bson:"native"`
	UserPreferred string `bson:"user_preferred"`
}

type MediaCoverImage struct {
	ExtraLarge string `bson:"image_extra_large"`
	Large      string `bson:"image_large"`
	Medium     string `bson:"image_medium"`
	Color      string `bson:"image_color"`
}

type MediaTrailer struct {
	ID        string `bson:"trailer_id"`
	Site      string `bson:"trailer_site"`
	Thumbnail string `bson:"trailer_thumbnail"`
}

type Mediatag struct {
	ID               int    `bson:"mediatag_id"`
	Name             string `bson:"mediatag_name"`
	Description      string `bson:"mediatag_description"`
	Category         string `bson:"mediatag_category"`
	Rank             int    `bson:"mediatag_rank"`
	IsGeneralSpoiler bool   `bson:"mediatag_is_general_spoiler"`
	IsMediaSpoiler   bool   `bson:"mediatag_is_media_spoiler"`
	IsAdult          bool   `bson:"mediatag_is_adult"`
	UserID           int    `bson:"mediatag_user_id"`
}

type Studio struct {
	ID   string `bson:"studio_id"`
	Name string `bson:"studio_name"`
}

type AnimeDocument struct {
	AniListID    int               `bson:"anilist_id"`
	MongoID      bson.ObjectID     `bson:"_id,omitempty"`
	Title        MediaTitle        `bson:"anime_title"`
	CoverImage   MediaCoverImage   `bson:"anime_cover_image"`
	Format       anime.MediaFormat `bson:"anime_format"`
	Status       anime.MediaStatus `bson:"anime_status"`
	Episodes     int               `bson:"anime_episodes"`
	Genres       []string          `bson:"anime_genres"`
	AverageScore int               `bson:"anime_average_score"`
	Popularity   int               `bson:"anime_popularity"`
	Trending     int               `bson:"anime_trending"`
	Description  string            `bson:"anime_description"`
	Trailer      MediaTrailer      `bson:"anime_trailer"`
	Studios      []Studio          `bson:"anime_studios"`
	Season       anime.MediaSeason `bson:"anime_season"`
	SeasonYear   int               `bson:"anime_season_year"`
	SeasonInt    int               `bson:"anime_season_int"`
	Tags         []Mediatag        `bson:"anime_tags"`
	SyncedAT     time.Time         `bson:"anime_syncedAT"`
}

type FuzzyDateDocument struct {
	Year  int `bson:"fuzzy_date_year"`
	Month int `bson:"fuzzy_date_month"`
	Day   int `bson:"fuzzy_date_day"`
}

type ListDocument struct {
	AniListID   int                  `bson:"anilist_id"`
	MongoID     bson.ObjectID        `bson:"_id,omitempty"`
	MongoUserID bson.ObjectID        `bson:"user_id"`
	Status      list.MediaListStatus `bson:"list_status"`
	Score       float64              `bson:"list_score"`
	Progress    int                  `bson:"list_progress"`
	Repeat      int                  `bson:"list_repeat"`
	Priority    int                  `bson:"list_priority"`
	Private     *bool                `bson:"list_private"`
	Notes       string               `bson:"list_notes"`
	StartedAt   FuzzyDateDocument    `bson:"list_startedAt"`
	FinishedAt  FuzzyDateDocument    `bson:"list_finishedAt"`
	CreatedAt   time.Time            `bson:"list_createdAt"`
	UpdatedAt   time.Time            `bson:"list_updatedAt"`
}

type UserAvatarDocument struct {
	Large  string `bson:"user_avatar_large"`
	Medium string `bson:"user_avatar_medium"`
}
type UserDocument struct {
	MongoUserID bson.ObjectID      `bson:"_id,omitempty"`
	Name        string             `bson:"user_name"`
	Password    string             `bson:"user_password"`
	About       string             `bson:"user_about"`
	Avatar      UserAvatarDocument `bson:"user_avatar"`
	BannerImage string             `bson:"user_banner_image"`
	Email       string             `bson:"user_email"`
	CreatedAt   time.Time          `bson:"user_created_at"`
	UpdatedAt   time.Time          `bson:"user_updated_at"`
}

func (d *UserDocument) ToDomainUser() *user.User {
	return &user.User{
		MongoUserID: d.MongoUserID,
		Name:        d.Name,
		Password:    d.Password,
		About:       d.About,
		Avatar: user.UserAvatar{
			Large:  d.Avatar.Large,
			Medium: d.Avatar.Medium,
		},
		Email:     d.Email,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func FromDomainList(d list.List, userID bson.ObjectID) *ListDocument {
	return &ListDocument{
		AniListID:   d.AniListID,
		MongoID:     d.MongoID,
		MongoUserID: userID,
		Status:      d.Status,
		Score:       d.Score,
		Progress:    d.Progress,
		Repeat:      d.Repeat,
		Priority:    d.Priority,
		Private:     d.Private,
		Notes:       d.Notes,
		StartedAt: FuzzyDateDocument{
			Year:  d.StartedAt.Year,
			Month: d.StartedAt.Month,
			Day:   d.StartedAt.Day,
		},
		FinishedAt: FuzzyDateDocument{
			Year:  d.FinishedAt.Year,
			Month: d.FinishedAt.Month,
			Day:   d.FinishedAt.Day,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (d *ListDocument) ToDomainList() *list.List {
	return &list.List{
		AniListID:   d.AniListID,
		MongoID:     d.MongoID,
		MongoUserID: d.MongoUserID,
		Status:      d.Status,
		Score:       d.Score,
		Progress:    d.Progress,
		Repeat:      d.Repeat,
		Priority:    d.Priority,
		Private:     d.Private,
		Notes:       d.Notes,
		StartedAt: list.FuzzyDate{
			Year:  d.StartedAt.Year,
			Month: d.StartedAt.Month,
			Day:   d.StartedAt.Day,
		},
		FinishedAt: list.FuzzyDate{
			Year:  d.FinishedAt.Year,
			Month: d.FinishedAt.Month,
			Day:   d.FinishedAt.Day,
		},
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func (d *AnimeDocument) ToDomain() *anime.Anime {
	return &anime.Anime{
		AniListID: d.AniListID,
		MongoID:   d.MongoID,
		Title: anime.MediaTitle{
			Romaji:        d.Title.Romaji,
			English:       d.Title.English,
			Native:        d.Title.Native,
			UserPreferred: d.Title.UserPreferred,
		},
		CoverImage: anime.MediaCoverImage{
			ExtraLarge: d.CoverImage.ExtraLarge,
			Large:      d.CoverImage.Large,
			Medium:     d.CoverImage.Medium,
			Color:      d.CoverImage.Color,
		},
		Trailer: anime.MediaTrailer{
			ID:        d.Trailer.ID,
			Site:      d.Trailer.Site,
			Thumbnail: d.Trailer.Thumbnail,
		},
		Format:       d.Format,
		Status:       d.Status,
		Season:       d.Season,
		Episodes:     d.Episodes,
		Genres:       d.Genres,
		AverageScore: d.AverageScore,
		Popularity:   d.Popularity,
		Trending:     d.Trending,
		Description:  d.Description,
		SeasonYear:   d.SeasonYear,
		SeasonInt:    d.SeasonInt,
		SyncedAT:     d.SyncedAT,
		Studios:      ToStudiosDomain(d.Studios),
		Tags:         ToTagsDomain(d.Tags),
	}
}

func ToStudiosDomain(studios []Studio) []anime.Studio {
	result := make([]anime.Studio, len(studios))
	for i, s := range studios {
		result[i] = anime.Studio{ID: s.ID, Name: s.Name}
	}
	return result
}

func ToTagsDomain(tags []Mediatag) []anime.Mediatag {
	result := make([]anime.Mediatag, len(tags))
	for i, t := range tags {
		result[i] = anime.Mediatag{
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
	return result
}

func FromDomain(a *anime.Anime) *AnimeDocument {
	studios := make([]Studio, len(a.Studios))
	for i, s := range a.Studios {
		studios[i] = Studio{ID: s.ID, Name: s.Name}
	}

	tags := make([]Mediatag, len(a.Tags))
	for i, t := range a.Tags {
		tags[i] = Mediatag{
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

	return &AnimeDocument{
		AniListID: a.AniListID,
		MongoID:   a.MongoID,
		Title: MediaTitle{
			Romaji:        a.Title.Romaji,
			English:       a.Title.English,
			Native:        a.Title.Native,
			UserPreferred: a.Title.UserPreferred,
		},
		CoverImage: MediaCoverImage{
			ExtraLarge: a.CoverImage.ExtraLarge,
			Large:      a.CoverImage.Large,
			Medium:     a.CoverImage.Medium,
			Color:      a.CoverImage.Color,
		},
		Trailer: MediaTrailer{
			ID:        a.Trailer.ID,
			Site:      a.Trailer.Site,
			Thumbnail: a.Trailer.Thumbnail,
		},
		Format:       a.Format,
		Status:       a.Status,
		Season:       a.Season,
		Episodes:     a.Episodes,
		Genres:       a.Genres,
		AverageScore: a.AverageScore,
		Popularity:   a.Popularity,
		Trending:     a.Trending,
		Description:  a.Description,
		SeasonYear:   a.SeasonYear,
		SeasonInt:    a.SeasonInt,
		Tags:         tags,
		Studios:      studios,
		SyncedAT:     a.SyncedAT,
	}
}
