package main

import (
	"log"
	"net/http"
	"os"

	"github.com/BerkBugur/Go-Project/controllers"
	"github.com/BerkBugur/Go-Project/docs"
	"github.com/BerkBugur/Go-Project/initializers"
	"github.com/BerkBugur/Go-Project/models"

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

func init() {
	prometheus.MustRegister(requestCount)
	initializers.LoadEnvVars()
	initializers.ConnectDB()
	initializers.DB.AutoMigrate(&models.Task{})

	// Log dosyasını aç
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		// Log dosyasını ayarla
		logrus.SetOutput(file)
		logrus.SetFormatter(&logrus.TextFormatter{})
	} else {
		logrus.Info("Failed to log to file, using default stderr")
	}
}

func main() {
	// Prometheus
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Fatal(http.ListenAndServe(":8081", nil))
	}()

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

	// Logrus test
	logrus.Info("Uygulama başladı")
	logrus.Warn("Bu bir uyarı mesajıdır")
	logrus.Error("Bu bir hata mesajıdır")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
