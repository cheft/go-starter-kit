package controllers

import (
	"net/http"
	"strconv"

	models "go-api/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	model models.User
	path  string
}

func NewUserController(router *gin.Engine) {
	p := models.NewUser()
	ctrl := UserController{path: "users", model: p}
	router.GET(ctrl.path, ctrl.Find)
	router.POST(ctrl.path, ctrl.Create)
	router.GET(ctrl.path+"/:id", ctrl.First)
	router.PUT(ctrl.path+"/:id", ctrl.Update)
	router.DELETE(ctrl.path+"/:id", ctrl.Delete)
}

func (ctrl UserController) Find(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	result, count := ctrl.model.Find(page, pageSize)
	pages := count / pageSize
	if count%pageSize != 0 {
		pages++
	}
	c.JSON(http.StatusOK, gin.H{
		"list":     result,
		"page":     page,
		"pageSize": pageSize,
		"count":    count,
		"pages":    pages,
	})
}

func (ctrl UserController) Create(c *gin.Context) {
	var u models.User
	c.Bind(&u)
	ctrl.model.Create(&u)
	c.JSON(http.StatusOK, u)
}

func (ctrl UserController) First(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result := ctrl.model.First(id)
	c.JSON(http.StatusOK, result.Value)
}

func (ctrl UserController) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var u models.User
	c.Bind(&u)
	u.ID = uint(id)
	ctrl.model.Update(&u)
	c.JSON(http.StatusOK, u)
}

func (ctrl UserController) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	list := ctrl.model.Delete(uint(id))
	c.JSON(http.StatusOK, list.Value)
}
