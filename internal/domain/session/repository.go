package session

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Session interface {
	CreateSession(ctx context.Context, userID bson.ObjectID) (string, error)
	GetSession(ctx context.Context, sessionID string) (bson.ObjectID, error)
	DeleteSession(ctx context.Context, sessionID string) error
}
