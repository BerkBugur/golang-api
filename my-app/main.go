package main

import (
	"log"
	"net/http"

	"github.com/BerkBugur/Go-Project/controllers"
	"github.com/BerkBugur/Go-Project/docs"
	"github.com/BerkBugur/Go-Project/initializers"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method"},
	)
)

// @title Documenting API
// @version 1
// @Description CRUD Operations

func init() {
	prometheus.MustRegister(requestCount)
	initializers.LoadEnvVars()
	initializers.ConnectDB()
}

func main() {
	// Logrus
	logrus.SetFormatter(&logrus.TextFormatter{})
	//Prometheus
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
	//TESTS
	logrus.Info("Uygulama başladı")
	logrus.Warn("Bu bir uyarı mesajıdır")
	logrus.Error("Bu bir hata mesajıdır")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080

}
