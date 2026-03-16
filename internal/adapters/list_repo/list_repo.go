package list_repo

import (
	"Katara/internal/adapters/documents"
	"Katara/internal/domain/list"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ListsRepo struct {
	db *mongo.Collection
}

func NewListRepository(db *mongo.Database) *ListsRepo {
	return &ListsRepo{
		db: db.Collection("lists"),
	}
}

func (s *ListsRepo) GetAllLists(ctx context.Context, userID bson.ObjectID, mediaListStatus []list.MediaListStatus) ([]list.List, error) {

	var filter bson.M

	filter = bson.M{"user_id": userID}

	if len(mediaListStatus) > 0 && len(mediaListStatus) < 3 {
		filter["list_status"] = bson.M{"$in": mediaListStatus}
	}

	cursor, err := s.db.Find(ctx, filter, nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var docs []documents.ListDocument
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}

	result := make([]list.List, len(docs))

	for i, v := range docs {
		result[i] = *v.ToDomainList()
	}

	return result, err
}

func (s *ListsRepo) GetListByStatus(ctx context.Context, userID bson.ObjectID, mediaListStatus list.MediaListStatus) ([]list.List, error) {
	cursor, err := s.db.Find(ctx, bson.M{"list_user_id": userID, "list_status": mediaListStatus})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var docs []documents.ListDocument
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}

	result := make([]list.List, len(docs))

	for i, v := range docs {
		result[i] = *v.ToDomainList()
	}

	return result, err
}

func (s *ListsRepo) AddList(ctx context.Context, MongoUserID bson.ObjectID, AniListID int, mediaListStatus list.MediaListStatus, Score float64, Progress int, Repeat int, Private *bool, Notes string, StartedAt documents.FuzzyDateDocument, FinishedAt documents.FuzzyDateDocument) error {
	doc := documents.ListDocument{
		MongoUserID: MongoUserID,
		AniListID:   AniListID,
		Status:      mediaListStatus,
		Score:       Score,
		Progress:    Progress,
		Repeat:      Repeat,
		Private:     Private,
		Notes:       Notes,
		StartedAt:   StartedAt,
		FinishedAt:  FinishedAt,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err := s.db.InsertOne(ctx, doc)

	if err != nil {
		return err
	}

	return err
}

func (s *ListsRepo) UpdateList(ctx context.Context, userID bson.ObjectID, list list.List) error {
	filter := bson.M{
		"_id":     list.MongoID,
		"user_id": userID,
	}
	update := bson.M{}
	if list.AniListID != 0 {
		update["anilist_id"] = list.AniListID
	}
	if list.Status != "" {
		update["list_status"] = list.Status
	}
	if list.Score != 0 {
		update["list_score"] = list.Score
	}
	if list.Progress != 0 {
		update["list_progress"] = list.Progress
	}
	if list.Repeat != 0 {
		update["list_repeat"] = list.Repeat
	}
	if list.Priority != 0 {
		update["list_priority"] = list.Priority
	}
	if list.Private != nil {
		update["list_private"] = list.Private
	}
	if list.Notes != "" {
		update["list_notes"] = list.Notes
	}
	if list.StartedAt.Year != 0 {
		update["list_startedAt"] = list.StartedAt
	}
	if list.FinishedAt.Year != 0 {
		update["list_finishedAt"] = list.FinishedAt
	}
	update["list_updatedAt"] = time.Now()

	_, err := s.db.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return err
	}

	return nil
}

func (s *ListsRepo) RemoveList(ctx context.Context, userID bson.ObjectID, AniListID int) error {
	filter := bson.M{
		"anilist_id": AniListID,
		"user_id":    userID,
	}

	_, err := s.db.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}
