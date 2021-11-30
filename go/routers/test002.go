package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func test002Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Test002 World!!!",
	})
}
func Test002(e *gin.Engine)  {
	e.GET("/test002", test002Handler)
}
