package main

import (
	"github.com/ergildo/go-lang-rest-app-mysql-example/model"
	"github.com/ergildo/go-lang-rest-app-mysql-example/service"
	"github.com/ergildo/go-lang-rest-app-mysql-example/setup"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	setup.SetUpDB()
	router := gin.Default()
	router.POST("/", save)
	router.PUT("/", update)
	router.GET("/", listAll)
	router.GET("/:id", findById)
	router.DELETE("/:id", delete)
	router.Run(":8080")

}

func save(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	userSaved := service.Save(user)

	c.JSON(http.StatusOK, userSaved)

}

func update(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}
	userSaved := service.Update(user)

	c.JSON(http.StatusOK, userSaved)
}

func listAll(c *gin.Context) {
	users := service.ListAll()
	c.JSON(http.StatusOK, users)
}

func findById(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.ParseInt(paramId, 10, 64)

	user := service.FindById(id)
	c.JSON(http.StatusOK, user)
}

func delete(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.ParseInt(paramId, 10, 64)
	service.Delete(id)
	c.Status(http.StatusOK)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
