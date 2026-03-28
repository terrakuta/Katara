package middleware

import (
	"Katara/internal/adapters/redis"
	"context"

	"github.com/gin-gonic/gin"
)

type contextKey string

const MongoIDkey contextKey = "mongoID"

func AuthMiddleware(repo *redis.SessionRepo) gin.HandlerFunc {
	return func(c *gin.Context) {

		sessionID, err := c.Cookie("session_id")

		if err != nil {
			c.Next()
			return
		}

		mongoID, err := repo.GetSession(c, sessionID)
		if err != nil {
			c.Next()
			return
		}

		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), MongoIDkey, mongoID))
		c.Next()
	}
}
