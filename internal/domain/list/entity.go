package list

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MediaListStatus string

const (
	StatusCURRENT   MediaListStatus = "CURRENT"
	StatusPLANNING  MediaListStatus = "PLANNING"
	StatusCOMPLETED MediaListStatus = "COMPLETED"
	StatusDROPPED   MediaListStatus = "DROPPED"
	StatusPAUSED    MediaListStatus = "PAUSED"
	StatusREPEATING MediaListStatus = "REPEATING"
)

type FuzzyDate struct {
	Year  int
	Month int
	Day   int
}

type List struct {
	AniListID   int
	MongoID     bson.ObjectID
	MongoUserID bson.ObjectID
	Status      MediaListStatus
	Score       float64
	Progress    int
	Repeat      int
	Priority    int
	Private     *bool
	Notes       string
	StartedAt   FuzzyDate
	FinishedAt  FuzzyDate
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
