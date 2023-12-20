package main

import (
	"log"
	"net/http"
	"os"

	"github.com/BerkBugur/Go-Project/controllers"
	"github.com/BerkBugur/Go-Project/docs"
	"github.com/BerkBugur/Go-Project/initializers"
	"github.com/BerkBugur/Go-Project/middleware"

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
	gin.SetMode(gin.ReleaseMode)

	// Prometheus
	prometheus.MustRegister(requestCount)
	initializers.LoadEnvVars()
	initializers.ConnectDB()
	// Sync database models
	initializers.SyncDatabase()
	// Open log file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		// Log dosyasını ayarla
		logrus.SetOutput(file)
		logrus.SetFormatter(&logrus.TextFormatter{})
	} else {
		logrus.Info("Failed to log to file, using default stderr")
	}
}

//@securityDefinitions.apikey jwt
//@in header
//@name Authorization

func main() {
	// CompileDaemon -command="./main.go"
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
		tasks.GET("/paged", middleware.RequireAuth, controllers.GetAllTaskByPage)

	}
	users := r.Group("/users")
	{
		users.POST("/signup", controllers.SignUp)
		users.POST("/login", controllers.Login)
		users.GET("/validate", middleware.RequireAuth, controllers.Validate) // Require Auth
	}

	// Logrus test
	logrus.Info("App Started.")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
