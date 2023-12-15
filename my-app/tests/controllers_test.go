package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BerkBugur/Go-Project/controllers"
	"github.com/BerkBugur/Go-Project/initializers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTaskShowByID(t *testing.T) {
	// Test ortamını kur
	initializers.ConnectDB()
	initializers.LoadEnvVars()

	// Create a test router
	router := gin.Default()

	// Add the TaskShowByID route
	router.GET("/tasks/:id", controllers.TaskShowByID)

	// Prepare a request with a specific task ID
	req, err := http.NewRequest("GET", "/tasks/1", nil)
	assert.NoError(t, err)

	// Create a test response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// TODO: Add more assertions for the structure of the returned task
	// ...
}
