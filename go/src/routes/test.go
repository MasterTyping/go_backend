package routes

import "github.com/gin-gonic/gin"

const test = "test"

func Test(r *gin.Engine) {
	r.GET(test, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
}
