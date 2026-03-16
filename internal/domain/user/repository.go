package user

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id bson.ObjectID) (*User, error)
	CreateUser(ctx context.Context, name string, email string, password string) (*User, error)
	UpdateAvatar(ctx context.Context, mongoID bson.ObjectID, newAvatar UserAvatar) (*User, error)
	UpdateEmail(ctx context.Context, mongoID bson.ObjectID, newEmail string) (*User, error)
	UpdatePassword(ctx context.Context, mongoID bson.ObjectID, hashedPassword string) (*User, error)
	UpdateBannerImage(ctx context.Context, mongoID bson.ObjectID, newBannerImage string) (*User, error)
	UpdateAbout(ctx context.Context, mongoID bson.ObjectID, newAbout string) (*User, error)
}
