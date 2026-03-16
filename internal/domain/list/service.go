package list

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ListService struct {
	repo ListRepository
}

func (s *ListService) AddList(ctx context.Context, MongoUserID bson.ObjectID, AniListID int, mediaListStatus MediaListStatus, Score float64, Progress int, Repeat int, Private bool, Notes string, StartedAt FuzzyDate, FinishedAt FuzzyDate) error {
	if err := s.repo.AddList(ctx, MongoUserID, AniListID, mediaListStatus, Score, Progress, Repeat, Private, Notes, StartedAt, FinishedAt); err != nil {
		return err
	}

	return nil
}

func (s *ListService) UpdateList(ctx context.Context, userID bson.ObjectID, list List) error {

	if err := s.repo.UpdateList(ctx, userID, list); err != nil {
		return err
	}

	return nil
}

func (s *ListService) RemoveList(ctx context.Context, userID bson.ObjectID, AniListID int) error {

	if err := s.repo.RemoveList(ctx, userID, AniListID); err != nil {
		return err
	}

	return nil
}

func (s *ListService) GetAllLists(ctx context.Context, userID bson.ObjectID, mediaListStatus []MediaListStatus) ([]List, error) {

	list, err := s.repo.GetAllLists(ctx, userID, mediaListStatus)

	if err != nil {
		return nil, err
	}

	return list, err
}

func (s *ListService) GetListByStatus(ctx context.Context, userID bson.ObjectID, mediaListStatus MediaListStatus) ([]List, error) {

	list, err := s.repo.GetListByStatus(ctx, userID, mediaListStatus)

	if err != nil {
		return nil, err
	}

	return list, err
}
