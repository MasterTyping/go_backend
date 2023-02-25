package routes

import "github.com/gin-gonic/gin"

func InitializeDefault(r *gin.Engine) {
	r.Handle("GET", "/healthcheck", HealthCheck)

}

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{})
}
