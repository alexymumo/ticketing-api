package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func RateLimiter(redisClient *redis.Client, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		clientIp := c.ClientIP()
		key := "rate_limit:" + clientIp

		currentCount, err := redisClient.Incr(ctx, key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			c.Abort()
			return
		}
		if currentCount == 1 {
			redisClient.Expire(ctx, key, window)
		}
		if currentCount > int64(limit) {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}
		c.Next()
	}
}
