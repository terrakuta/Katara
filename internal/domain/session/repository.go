package session

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session interface {
	CreateSession(ctx context.Context, userID primitive.ObjectID) (string, error)
	GetSession(ctx context.Context, sessionID string) (primitive.ObjectID, error)
	DeleteSession(ctx context.Context, sessionID string) error
}
