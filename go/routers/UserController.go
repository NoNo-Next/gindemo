package routers

import (
	service2 "gindemo/go/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAll(c *gin.Context) {
	todoList, err := service2.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {

		c.HTML(http.StatusOK, "user/user.html", gin.H{
			"code": 200,
			"msg":  "success",
			"data": todoList,
		})
	}

}
func UserController(e *gin.Engine) {
	e.GET("/getAllUser", getAll)
}
