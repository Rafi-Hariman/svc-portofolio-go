package middleware

import "github.com/gin-gonic/gin"

const (
	PUBLIC  string = "public"
	PRIVATE string = "private"
)

func PublicMiddleware(c *gin.Context) {
	c.Set("requestSource", PUBLIC)
	c.Next()
}

func PrivateMiddleware(c *gin.Context) {
	c.Set("requestSource", PRIVATE)
	c.Next()
}
