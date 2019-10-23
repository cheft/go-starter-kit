package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeptController struct {
	path string
}

func NewDeptController(router *gin.Engine) {
	ctrl := DeptController{"depts"}
	router.GET(ctrl.path, ctrl.GET)
	router.POST(ctrl.path, ctrl.GET)
	router.GET(ctrl.path+"/:id", ctrl.GET)
	router.PUT(ctrl.path+"/:id", ctrl.GET)
	router.DELETE(ctrl.path+"/:id", ctrl.GET)
}

func (ctrl DeptController) GET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "hello"})
}
