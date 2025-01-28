package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var redisClient *redis.Client

func RateLimiter(limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorizaion")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		key := "rate_limit:" + userId.(string)

		currentCount, err := redisClient.Get(ctx, key).Int()
		if err == redis.Nil {
			redisClient.Set(ctx, key, 1, window)
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			c.Abort()
			return
		} else {
			if currentCount >= limit {
				c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
				c.Abort()
				return
			}
			redisClient.Incr(ctx, key)
		}
		c.Next()
	}
}
