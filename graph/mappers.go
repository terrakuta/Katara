package graph

import (
	"Katara/graph/model"
	"Katara/internal/domain/anime"
	"Katara/internal/domain/list"
	"Katara/internal/domain/user"
)

func toGraphqlUser(d *user.User) *model.User {
	return &model.User{
		Name:  d.Name,
		Email: d.Email,
		About: &d.About,
		Avatar: &model.UserAvatar{
			Large:  &d.Avatar.Large,
			Medium: &d.Avatar.Medium,
		},
		BannerImage: &d.BannerImage,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}
func fromGraphqlUserAvatar(d *model.UserAvatarInput) user.UserAvatar {
	result := user.UserAvatar{}
	if d.Large != nil {
		result.Large = *d.Large
	}
	if d.Medium != nil {
		result.Medium = *d.Medium
	}
	return result
}

func fromAddListInput(d model.AddListInput) (int, list.MediaListStatus, float64, int, int, bool, string, list.FuzzyDate, list.FuzzyDate) {
	var score float64
	if d.Score != nil {
		score = *d.Score
	}
	var progress int
	if d.Progress != nil {
		progress = int(*d.Progress)
	}
	var repeat int
	if d.Repeat != nil {
		repeat = int(*d.Repeat)
	}
	var private bool
	if d.Private != nil {
		private = *d.Private
	}
	var notes string
	if d.Notes != nil {
		notes = *d.Notes
	}
	var startedAt list.FuzzyDate
	if d.StartedAt != nil {
		startedAt = list.FuzzyDate{Year: int(d.StartedAt.Year), Month: int(d.StartedAt.Month), Day: int(d.StartedAt.Day)}
	}
	var finishedAt list.FuzzyDate
	if d.FinishedAt != nil {
		finishedAt = list.FuzzyDate{Year: int(d.FinishedAt.Year), Month: int(d.FinishedAt.Month), Day: int(d.FinishedAt.Day)}
	}
	return int(d.AnilistID), list.MediaListStatus(d.MediaListStatus), score, progress, repeat, private, notes, startedAt, finishedAt
}

func toGraphqlAnime(d *anime.Anime) *model.Anime {
	episodes := int32(d.Episodes)
	averageScore := int32(d.AverageScore)
	popularity := int32(d.Popularity)
	trending := int32(d.Trending)
	seasonYear := int32(d.SeasonYear)
	seasonInt := int32(d.SeasonInt)
	format := model.MediaFormat(d.Format)
	status := model.MediaStatus(d.Status)
	season := model.MediaSeason(d.Season)

	return &model.Anime{
		AniListID: int32(d.AniListID),
		MongoID:   d.MongoID.Hex(),
		Title: &model.MediaTitle{
			Romaji:        d.Title.Romaji,
			English:       d.Title.English,
			Native:        d.Title.Native,
			UserPreferred: d.Title.UserPreferred,
		},
		CoverImage: &model.MediaCoverImage{
			ExtraLarge: d.CoverImage.ExtraLarge,
			Large:      d.CoverImage.Large,
			Medium:     d.CoverImage.Medium,
			Color:      d.CoverImage.Color,
		},
		Trailer: &model.MediaTrailer{
			ID:        d.Trailer.ID,
			Site:      d.Trailer.Site,
			Thumbnail: d.Trailer.Thumbnail,
		},
		Format:       &format,
		Status:       &status,
		Season:       &season,
		Episodes:     &episodes,
		Genres:       d.Genres,
		AverageScore: &averageScore,
		Popularity:   &popularity,
		Trending:     &trending,
		Description:  &d.Description,
		SeasonYear:   &seasonYear,
		SeasonInt:    &seasonInt,
		SyncedAt:     &d.SyncedAT,
		Studios:      toStudiosModel(d.Studios),
		Tags:         toTagsModel(d.Tags),
	}
}
func toStudiosModel(studios []anime.Studio) []*model.Studio {
	result := make([]*model.Studio, len(studios))
	for i, s := range studios {
		result[i] = &model.Studio{ID: s.ID, Name: s.Name}
	}
	return result
}

func toTagsModel(tags []anime.Mediatag) []*model.Mediatag {
	result := make([]*model.Mediatag, len(tags))
	for i, t := range tags {
		result[i] = &model.Mediatag{
			ID:               int32(t.ID),
			Name:             t.Name,
			Description:      t.Description,
			Category:         t.Category,
			Rank:             int32(t.Rank),
			IsGeneralSpoiler: t.IsGeneralSpoiler,
			IsMediaSpoiler:   t.IsMediaSpoiler,
			IsAdult:          t.IsAdult,
			UserID:           int32(t.UserID),
		}
	}
	return result
}

func toAnimeFilter(d model.AnimeFilter) anime.AnimeFilter {
	filter := anime.AnimeFilter{}

	if d.AnilistID != nil {
		filter.AnilistID = int(*d.AnilistID)
	}
	if d.Sort != nil {
		filter.Sort = anime.MediaSort(*d.Sort)
	}
	if d.Status != nil {
		filter.Status = anime.MediaStatus(*d.Status)
	}
	if d.Season != nil {
		filter.Season = anime.MediaSeason(*d.Season)
	}
	if d.Format != nil {
		filter.Format = anime.MediaFormat(*d.Format)
	}
	if d.Year != nil {
		filter.Year = int(*d.Year)
	}
	if d.Page != nil {
		filter.Page = int(*d.Page)
	}
	if d.PerPage != nil {
		filter.PerPage = int(*d.PerPage)
	}
	for _, g := range d.Genre {
		if g != nil {
			filter.Genre = append(filter.Genre, *g)
		}
	}
	return filter
}
