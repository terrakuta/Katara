package user

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserAvatar struct {
	Large  string
	Medium string
}

type User struct {
	MongoUserID bson.ObjectID
	Name        string
	Password    string `json:"-"`
	About       string
	Avatar      UserAvatar
	BannerImage string
	Email       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
