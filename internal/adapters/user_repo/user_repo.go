package user_repo

import (
	"Katara/internal/adapters/documents"
	"Katara/internal/domain/anime"
	"Katara/internal/domain/user"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var UserNotFound = errors.New("not found")

type UserRepo struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepo {
	return &UserRepo{
		db: db.Collection("users"),
	}
}

func (s *UserRepo) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	var user documents.UserDocument
	err := s.db.FindOne(ctx, bson.M{"user_email": email}).Decode(&user)

	if err != nil {
		if err == anime.ErrNotFound {
			return nil, err
		}
		return nil, err
	}

	convertedUser := user.ToDomainUser()

	return convertedUser, err

}

func (s *UserRepo) GetUserByID(ctx context.Context, id primitive.ObjectID) (*user.User, error) {
	var user documents.UserDocument
	err := s.db.FindOne(ctx, bson.M{"_id": id}).Decode(&user)

	if err != nil {
		if err == anime.ErrNotFound {
			return nil, err
		}
		return nil, err
	}

	convertedUser := user.ToDomainUser()

	return convertedUser, err

}

func (s *UserRepo) CreateUser(ctx context.Context, name string, email string, password string) (*user.User, error) {
	doc := documents.UserDocument{
		Name:     name,
		Email:    email,
		Password: password,
	}
	_, err := s.db.InsertOne(ctx, doc)

	if err != nil {
		return nil, err
	}

	converted := doc.ToDomainUser()

	return converted, nil
}

func (s *UserRepo) UpdateAvatar(ctx context.Context, mongoID primitive.ObjectID, newAvatar user.UserAvatar) (*user.User, error) {
	update := bson.M{"$set": bson.M{"user_avatar": newAvatar}}
	_, err := s.db.UpdateOne(ctx, bson.M{"_id": mongoID}, update)

	if err != nil {
		if err == UserNotFound {
			return nil, err
		}
		return nil, err
	}

	var doc documents.UserDocument

	found := s.db.FindOne(ctx, bson.M{"_id": mongoID}).Decode(&doc)

	if found != nil {
		if found == UserNotFound {
			return nil, err
		}
		return nil, err
	}

	convertedUser := doc.ToDomainUser()

	return convertedUser, err
}

func (s *UserRepo) UpdateEmail(ctx context.Context, mongoID primitive.ObjectID, newEmail string) (*user.User, error) {
	update := bson.M{"$set": bson.M{"user_email": newEmail}}
	_, err := s.db.UpdateOne(ctx, bson.M{"_id": mongoID}, update)

	if err != nil {
		if err == UserNotFound {
			return nil, err
		}
		return nil, err
	}

	var doc documents.UserDocument

	found := s.db.FindOne(ctx, bson.M{"_id": mongoID}).Decode(&doc)

	if found != nil {
		if found == UserNotFound {
			return nil, err
		}
		return nil, err
	}

	convertedUser := doc.ToDomainUser()

	return convertedUser, err

}

func (s *UserRepo) UpdatePassword(ctx context.Context, mongoID primitive.ObjectID, hashedPassword string) (*user.User, error) {
	update := bson.M{"$set": bson.M{"user_password": hashedPassword}}
	_, err := s.db.UpdateOne(ctx, bson.M{"_id": mongoID}, update)

	if err != nil {
		if err == UserNotFound {
			return nil, err
		}
		return nil, err
	}

	var doc documents.UserDocument

	found := s.db.FindOne(ctx, bson.M{"_id": mongoID}).Decode(&doc)

	if found != nil {
		if found == UserNotFound {
			return nil, err
		}
		return nil, err
	}

	convertedUser := doc.ToDomainUser()

	return convertedUser, err
}

func (s *UserRepo) UpdateBannerImage(ctx context.Context, mongoID primitive.ObjectID, newBannerImage string) (*user.User, error) {
	update := bson.M{"$set": bson.M{"user_banner_image": newBannerImage}}
	_, err := s.db.UpdateOne(ctx, bson.M{"_id": mongoID}, update)

	if err != nil {
		if err == UserNotFound {
			return nil, err
		}
		return nil, err
	}

	var doc documents.UserDocument

	found := s.db.FindOne(ctx, bson.M{"_id": mongoID}).Decode(&doc)

	if found != nil {
		if found == UserNotFound {
			return nil, err
		}
		return nil, err
	}

	convertedUser := doc.ToDomainUser()

	return convertedUser, err
}

func (s *UserRepo) UpdateAbout(ctx context.Context, mongoID primitive.ObjectID, newAbout string) (*user.User, error) {
	update := bson.M{"$set": bson.M{"user_about": newAbout}}
	_, err := s.db.UpdateOne(ctx, bson.M{"_id": mongoID}, update)

	if err != nil {
		if err == UserNotFound {
			return nil, err
		}
		return nil, err
	}

	var doc documents.UserDocument

	found := s.db.FindOne(ctx, bson.M{"_id": mongoID}).Decode(&doc)

	if found != nil {
		if found == UserNotFound {
			return nil, err
		}
		return nil, err
	}

	convertedUser := doc.ToDomainUser()

	return convertedUser, err
}
