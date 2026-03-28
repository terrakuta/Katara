package main

import (
	"Katara/graph"
	"Katara/internal/adapters/anilist"
	"Katara/internal/adapters/anime_repo"
	"Katara/internal/adapters/list_repo"
	"Katara/internal/adapters/redis"
	"Katara/internal/adapters/user_repo"
	"Katara/internal/config"
	"Katara/internal/database"
	"Katara/internal/domain/anime"
	"Katara/internal/domain/list"
	"Katara/internal/domain/user"
	"Katara/internal/worker"
	"Katara/middleware"
	"context"
	"fmt"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("Cannot load env config")
	}

	db, err := database.MongoLoad(cfg)

	if err != nil {
		log.Fatalf("Cannot connect database")
	}

	redisDB, err := database.RedisLoad(cfg)
	if err != nil {
		fmt.Printf("failed to connect to redis server: %s\n", err.Error())
	}

	defer func() {
		if err := db.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	//------------------------------------ REPO / SERVICES

	animeRepo := anime_repo.NewAnimeRepository(db.Database(cfg.MONGO_DB))
	if animeRepo == nil {
		log.Fatalf("animeRepo is nil")
	}

	animeService := anime.NewAnimeService(animeRepo)

	userRepo := user_repo.NewUserRepository(db.Database(cfg.MONGO_DB))
	if userRepo == nil {
		log.Fatalf("userRepo is nil")
	}

	sessionRepo := redis.NewSessionRepoRepository(redisDB)
	userService := user.NewUserService(userRepo, *sessionRepo)

	listRepo := list_repo.NewListRepository(db.Database(cfg.MONGO_DB))
	if listRepo == nil {
		log.Fatalf("userRepo is nil")
	}

	listService := list.NewListService(listRepo)

	anilistClient := anilist.NewAnilistClient()
	w := worker.NewSyncWorker(anilistClient, animeRepo)
	if cfg.SYNC_ENABLED {
		w.Start()
	}

	//------------------------------------ HANDLERS

	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{AnimeService: animeService, UserService: userService, ListService: listService},
	}))
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.Options{})

	log.Printf("animeRepo: %v", animeRepo)
	log.Printf("animeService: %v", animeService)
	log.Printf("listRepo: %v", listRepo)
	log.Printf("listService: %v", listService)

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware(sessionRepo))

	protected.POST("/query", func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "writer", c.Writer)
		c.Request = c.Request.WithContext(ctx)
		srv.ServeHTTP(c.Writer, c.Request)
	})
	r.GET("/playground", func(c *gin.Context) {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Writer, c.Request)
	})

	log.Fatal(r.Run(":8080"))

}
