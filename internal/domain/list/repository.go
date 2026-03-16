package list

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ListRepository interface {
	GetAllLists(ctx context.Context, userID bson.ObjectID, mediaListStatus []MediaListStatus) ([]List, error)
	GetListByStatus(ctx context.Context, userID bson.ObjectID, mediaListStatus MediaListStatus) ([]List, error)
	AddList(ctx context.Context, MongoUserID bson.ObjectID, AniListID int, mediaListStatus MediaListStatus, Score float64, Progress int, Repeat int, Private bool, Notes string, StartedAt FuzzyDate, FinishedAt FuzzyDate) error
	UpdateList(ctx context.Context, userID bson.ObjectID, list List) error
	RemoveList(ctx context.Context, userID bson.ObjectID, AniListID int) error
}
