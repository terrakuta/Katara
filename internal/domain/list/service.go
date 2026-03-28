package list

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ListService struct {
	repo ListRepository
}

func NewListService(repo ListRepository) *ListService {
	log.Printf("NewListService called, repo = %v", repo)
	return &ListService{repo: repo}
}

func (s *ListService) AddList(ctx context.Context, userID bson.ObjectID, params List) error {
	if err := s.repo.AddList(ctx, userID, params); err != nil {
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

func (s *ListService) RemoveList(ctx context.Context, userID bson.ObjectID, AniListID int) (bool, error) {

	removed, err := s.repo.RemoveList(ctx, userID, AniListID)
	if err != nil {
		return false, err
	}

	return removed, nil
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
