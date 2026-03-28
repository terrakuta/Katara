package user

import (
	"Katara/internal/adapters/redis"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo        UserRepository
	sessionRepo redis.SessionRepo
}

func NewUserService(repo UserRepository, sessionRepo redis.SessionRepo) *UserService {
	return &UserService{repo: repo, sessionRepo: sessionRepo}
}

func (s *UserService) Login(ctx context.Context, email string, password string) (*User, string, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, "", err
	}

	sessionID, err := s.sessionRepo.CreateSession(ctx, user.MongoUserID)
	if err != nil {
		log.Printf("CreateSession error: %v", err)
		return nil, "", err
	}
	log.Printf("Session created: %s", sessionID)

	if err != nil {
		return nil, sessionID, err
	}

	return user, sessionID, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id bson.ObjectID) (*User, error) {
	user, err := s.repo.GetUserByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Register(ctx context.Context, name string, email string, password string) (*User, error) {
	_, err := s.repo.GetUserByEmail(ctx, email)
	if err == nil {
		return nil, errors.New("user already exists")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user, err := s.repo.CreateUser(ctx, name, email, string(pass))
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (s *UserService) UpdateEmail(ctx context.Context, mongoID bson.ObjectID, newEmail string) (*User, error) {
	user, err := s.repo.UpdateEmail(ctx, mongoID, newEmail)

	if err != nil {
		return nil, err
	}

	return user, nil
}
func (s *UserService) UpdateAvatar(ctx context.Context, mongoID bson.ObjectID, newAvatar UserAvatar) (*User, error) {
	user, err := s.repo.UpdateAvatar(ctx, mongoID, newAvatar)

	if err != nil {
		return nil, err
	}

	return user, nil
}
func (s *UserService) UpdatePassword(ctx context.Context, mongoID bson.ObjectID, oldPassword string, newPassword string) (*User, error) {
	foundUser, err := s.repo.GetUserByID(ctx, mongoID)

	if err != nil {
		return nil, err
	}
	oldPass := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(oldPassword))

	if oldPass != nil {
		return nil, oldPass
	}

	newPass, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user, err := s.repo.UpdatePassword(ctx, mongoID, string(newPass))

	if err != nil {
		return nil, err
	}

	return user, nil
}
func (s *UserService) UpdateBannerImage(ctx context.Context, mongoID bson.ObjectID, newBannerImage string) (*User, error) {
	user, err := s.repo.UpdateBannerImage(ctx, mongoID, newBannerImage)

	if err != nil {
		return nil, err
	}

	return user, nil
}
func (s *UserService) UpdateAbout(ctx context.Context, mongoID bson.ObjectID, newAbout string) (*User, error) {
	user, err := s.repo.UpdateAbout(ctx, mongoID, newAbout)

	if err != nil {
		return nil, err
	}

	return user, nil
}
