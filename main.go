package main

import (
	controllers "go-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	controllers.NewUserController(router)
	controllers.NewDeptController(router)
	router.Run() // listen and serve on 0.0.0.0:8080
}
