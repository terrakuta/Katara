package anime

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MediaTitle struct {
	Romaji        string
	English       string
	Native        string
	UserPreferred string
}

type MediaCoverImage struct {
	ExtraLarge string
	Large      string
	Medium     string
	Color      string
}

type MediaTrailer struct {
	ID        string
	Site      string
	Thumbnail string
}

type Mediatag struct {
	ID               int
	Name             string
	Description      string
	Category         string
	Rank             int
	IsGeneralSpoiler bool
	IsMediaSpoiler   bool
	IsAdult          bool
	UserID           int
}

type MediaFormat string

const (
	FormatTV       MediaFormat = "TV"
	FormatTV_SHORT MediaFormat = "TV_SHORT"
	FormatMOVIE    MediaFormat = "MOVIE"
	FormatSPECIAL  MediaFormat = "SPECIAL"
	FormatOVA      MediaFormat = "OVA"
	FormatONA      MediaFormat = "ONA"
	FormatMUSIC    MediaFormat = "MUSIC"
	FormatMANGA    MediaFormat = "MANGA"
	FormatNOVEL    MediaFormat = "NOVEL"
	FormatONE_SHOT MediaFormat = "ONE_SHOT"
)

type MediaStatus string

const (
	StatusFINISHED         MediaStatus = "FINISHED"
	StatusRELEASED         MediaStatus = "RELEASING"
	StatusNOT_YET_RELEASED MediaStatus = "NOT_YET_RELEASED"
	StatusCANCELLED        MediaStatus = "CANCELLED"
	StatusHIATUS           MediaStatus = "HIATUS"
)

type MediaSeason string

const (
	SeasonWINTER MediaSeason = "WINTER"
	SeasonSPRING MediaSeason = "SPRING"
	SeasonSUMMER MediaSeason = "SUMMER"
	SeasonFALL   MediaSeason = "FALL"
)

type Studio struct {
	ID   string
	Name string
}

type MediaSort string

const (
	SortTitleRomaji      MediaSort = "TITLE_ROMAJI"
	SortTitleRomajiDesc  MediaSort = "TITLE_ROMAJI_DESC"
	SortTitleEnglish     MediaSort = "TITLE_ENGLISH"
	SortTitleEnglishDesc MediaSort = "TITLE_ENGLISH_DESC"
	SortTitleNative      MediaSort = "TITLE_NATIVE"
	SortTitleNativeDesc  MediaSort = "TITLE_NATIVE_DESC"
	SortFormat           MediaSort = "FORMAT"
	SortFormatDesc       MediaSort = "FORMAT_DESC"
	SortStartDate        MediaSort = "START_DATE"
	SortStartDateDesc    MediaSort = "START_DATE_DESC"
	SortEndDate          MediaSort = "END_DATE"
	SortEndDateDesc      MediaSort = "END_DATE_DESC"
	SortScore            MediaSort = "SCORE"
	SortScoreDesc        MediaSort = "SCORE_DESC"
	SortPopularity       MediaSort = "POPULARITY"
	SortPopularityDesc   MediaSort = "POPULARITY_DESC"
	SortTrending         MediaSort = "TRENDING"
	SortTrendingDesc     MediaSort = "TRENDING_DESC"
	SortEpisodes         MediaSort = "EPISODES"
	SortEpisodesDesc     MediaSort = "EPISODES_DESC"
	SortDuration         MediaSort = "DURATION"
	SortDurationDesc     MediaSort = "DURATION_DESC"
	SortStatus           MediaSort = "STATUS"
	SortStatusDesc       MediaSort = "STATUS_DESC"
	SortUpdatedAt        MediaSort = "UPDATED_AT"
	SortUpdatedAtDesc    MediaSort = "UPDATED_AT_DESC"
	SortSearchMatch      MediaSort = "SEARCH_MATCH"
)

type AnimeFilter struct {
	AnilistID int
	Sort      MediaSort
	Status    MediaStatus
	Season    MediaSeason
	Format    MediaFormat
	Genre     []string
	Year      int
	Page      int
	PerPage   int
}

type Anime struct {
	AniListID    int
	MongoID      bson.ObjectID
	Title        MediaTitle
	CoverImage   MediaCoverImage
	Format       MediaFormat
	Status       MediaStatus
	Episodes     int
	Genres       []string
	AverageScore int
	Popularity   int
	Trending     int
	Description  string
	Trailer      MediaTrailer
	Studios      []Studio
	Season       MediaSeason
	SeasonYear   int
	SeasonInt    int
	Tags         []Mediatag
	SyncedAT     time.Time
}
