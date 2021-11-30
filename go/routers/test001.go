package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func test001Handler(c *gin.Context) {
	param := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"message": "Test World!!!"+param,
	})
}
func Test001(e *gin.Engine)  {
	e.GET("/test001", test001Handler)
}
