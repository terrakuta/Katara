package graph

import (
	"Katara/graph/model"
	"Katara/internal/domain/anime"
	"Katara/internal/domain/list"
	"Katara/internal/domain/user"

	"go.mongodb.org/mongo-driver/v2/bson"
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

func MapDomainToGraphql(domainLists []list.List) []*model.List {
	result := make([]*model.List, len(domainLists))

	for i, d := range domainLists {
		score := d.Score
		progress := int32(d.Progress)
		repeat := int32(d.Repeat)
		priority := int32(d.Priority)
		notes := d.Notes

		private := false
		if d.Private != nil {
			private = *d.Private
		}

		result[i] = &model.List{
			AniListID: int32(d.AniListID),
			Status:    model.MediaListStatus(d.Status),
			Score:     &score,
			Progress:  &progress,
			Repeat:    &repeat,
			Priority:  &priority,
			Private:   private,
			Notes:     &notes,
			StartedAt: &model.FuzzyDate{
				Year:  int32(d.StartedAt.Year),
				Month: int32(d.StartedAt.Month),
				Day:   int32(d.StartedAt.Day),
			},
			FinishedAt: &model.FuzzyDate{
				Year:  int32(d.FinishedAt.Year),
				Month: int32(d.FinishedAt.Month),
				Day:   int32(d.FinishedAt.Day),
			},
		}
	}
	return result
}

func MapSliceDomainToGraphql(domainLists []list.List) []*model.List {
	result := make([]*model.List, len(domainLists))

	for i, d := range domainLists {
		score := d.Score
		progress := int32(d.Progress)
		repeat := int32(d.Repeat)
		priority := int32(d.Priority)
		notes := d.Notes

		private := false
		if d.Private != nil {
			private = *d.Private
		}

		result[i] = &model.List{
			AniListID: int32(d.AniListID),
			Status:    model.MediaListStatus(d.Status),
			Score:     &score,
			Progress:  &progress,
			Repeat:    &repeat,
			Priority:  &priority,
			Private:   private,
			Notes:     &notes,
			StartedAt: &model.FuzzyDate{
				Year:  int32(d.StartedAt.Year),
				Month: int32(d.StartedAt.Month),
				Day:   int32(d.StartedAt.Day),
			},
			FinishedAt: &model.FuzzyDate{
				Year:  int32(d.FinishedAt.Year),
				Month: int32(d.FinishedAt.Month),
				Day:   int32(d.FinishedAt.Day),
			},
		}
	}
	return result
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

func fromGraphqlAddListInput(d *model.AddListInput, mongoUserID bson.ObjectID) list.List {
	mapDate := func(input *model.FuzzyDateInput) list.FuzzyDate {
		if input == nil {
			return list.FuzzyDate{}
		}
		return list.FuzzyDate{
			Year:  int(input.Year),
			Month: int(input.Month),
			Day:   int(input.Day),
		}
	}

	result := list.List{
		AniListID:   int(d.AniListID),
		MongoUserID: mongoUserID,
		Status:      list.MediaListStatus(d.Status),
		Private:     &d.Private,
		StartedAt:   mapDate(d.StartedAt),
		FinishedAt:  mapDate(d.FinishedAt),
	}

	if d.Score != nil {
		result.Score = *d.Score
	}
	if d.Progress != nil {
		result.Progress = int(*d.Progress)
	}
	if d.Repeat != nil {
		result.Repeat = int(*d.Repeat)
	}
	if d.Priority != nil {
		result.Priority = int(*d.Priority)
	}
	if d.Notes != nil {
		result.Notes = *d.Notes
	}

	return result
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
