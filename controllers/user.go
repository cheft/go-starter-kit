package controllers

import (
	"fmt"
	"net/http"

	models "go-api/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	model models.Product
	path  string
}

func NewUserController(router *gin.Engine) {
	p := models.NewProduct()
	ctrl := UserController{path: "users", model: p}
	router.GET(ctrl.path, ctrl.Get)
	router.GET(ctrl.path+"/create", ctrl.Create)
}

func (ctrl UserController) Get(c *gin.Context) {
	list := ctrl.model.GetAll()
	fmt.Println(list)
	c.JSON(http.StatusOK, list.Value)
}

func (ctrl UserController) Create(c *gin.Context) {
	ctrl.model.Create()
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}
