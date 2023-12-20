package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BerkBugur/Go-Project/initializers"
	"github.com/BerkBugur/Go-Project/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	dsn := "host= user= password= dbname= port=5432 sslmode=disable" //  you can create db https://www.elephantsql.com/ for testing.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	// Veritabanı şemalarını burada başlatabilirsiniz
	db.AutoMigrate(&models.Task{})

	return db
}

func TestTaskCreate(t *testing.T) {
	testDB := setupTestDB()
	initializers.DB = testDB

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/tasks", TaskCreate)

	task := map[string]string{
		"title":       "Test Task",
		"description": "Test Description",
		"status":      "pending",
	}
	taskJSON, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Contains(t, response, "task")
}

func TestGetAllTask(t *testing.T) {
	testDB := setupTestDB()
	initializers.DB = testDB

	// Test ortamını ayarla
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/tasks", GetAllTask)

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "Status code should be 200")

	var response interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	t.Log(response)
}
