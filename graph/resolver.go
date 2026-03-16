package graph

import (
	"Katara/internal/domain/anime"
	"Katara/internal/domain/list"
	"Katara/internal/domain/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.
type Resolver struct {
	AnimeService *anime.AnimeService
	UserService  *user.UserService
	ListService  *list.ListService
}
