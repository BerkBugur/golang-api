package main

import (
	"log"
	"net/http"

	"github.com/BerkBugur/Go-Project/controllers"
	"github.com/BerkBugur/Go-Project/docs"
	"github.com/BerkBugur/Go-Project/initializers"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Documenting API
// @version 1
// @Description CRUD

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Fatal(http.ListenAndServe(":8081", nil))
	}()
	//CompileDaemon -command="./go-crud"
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	tasks := r.Group("/tasks")
	{
		tasks.GET("", controllers.GetAllTask)
		tasks.POST("", controllers.TaskCreate)
		tasks.GET("/:id", controllers.TaskShowByID)
		tasks.PUT("/:id", controllers.TaskUpdate)
		tasks.DELETE("/:id", controllers.TaskDelete)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080
}
