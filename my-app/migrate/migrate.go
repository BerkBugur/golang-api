package main

import (
	"github.com/BerkBugur/Go-Project/initializers"
	"github.com/BerkBugur/Go-Project/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Task{})
}
