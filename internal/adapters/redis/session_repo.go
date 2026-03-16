package redis

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionRepo struct {
	client *redis.Client
}

func NewSessionRepoRepository(client *redis.Client) *SessionRepo {
	return &SessionRepo{
		client: client,
	}
}

func (s *SessionRepo) CreateSession(ctx context.Context, userID primitive.ObjectID) (string, error) {
	sessionID := uuid.New().String()
	value := userID.Hex()

	err := s.client.Set(ctx, "session:"+sessionID, value, 30*time.Minute).Err()

	if err != nil {
		return "", err
	}
	return sessionID, nil
}

func (s *SessionRepo) GetSession(ctx context.Context, sessionID string) (primitive.ObjectID, error) {
	val, err := s.client.Get(ctx, "session:"+sessionID).Result()
	if err != nil {
		return primitive.ObjectID{}, err
	}

	userID, err := primitive.ObjectIDFromHex(val)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return userID, nil
}

func (s *SessionRepo) DeleteSession(ctx context.Context, sessionID string) error {
	err := s.client.Del(ctx, "session:"+sessionID).Err()

	if err != nil {
		return err
	}
	return nil
}
