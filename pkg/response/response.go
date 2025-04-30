package response

import "github.com/gin-gonic/gin"

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"error": message,
	})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"data": data,
	})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(201, gin.H{
		"data": data,
	})
}
