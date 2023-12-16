package controllers

// swag init --parseDependency true
import (
	"github.com/BerkBugur/Go-Project/initializers"
	"github.com/BerkBugur/Go-Project/models"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// @Title Users
// @Summary Create a task in db
// @Tags Tasks
// @Success      200  {object}  models.Task
// @Router       /tasks/ [post]
// @Param title formData string true "Task title"
// @Param description formData string true "Task description"
// @Param status formData string true "Task status"
func TaskCreate(c *gin.Context) {
	logrus.Info("TaskCreate endpoint")
	var body struct {
		Title       string `form:"title" binding:"required"`
		Description string `form:"description" binding:"required"`
		Status      string `form:"status" binding:"required"`
	}

	if err := c.ShouldBind(&body); err != nil {
		logrus.WithError(err).Error("Failed to bind request body")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{Title: body.Title, Description: body.Description, Status: body.Status}

	resultChan := make(chan error)
	pool.wg.Add(1)
	go func() {
		defer pool.wg.Done()
		pool.jobQueue <- TaskJob{Task: task, Result: resultChan}
	}()

	err := <-resultChan
	if err != nil {
		logrus.WithError(err).Error("Failed to process task creation job")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	logrus.Infof("Task created. Task ID: %d", task.ID)
	c.JSON(200, gin.H{"task": task})
}

// GetAllTask returns a JSON response with all tasks in the database.
// @Summary Get all tasks from the database
// @Tags Tasks
// @Produce json
// @Success 200 {array} models.Task
// @Router /tasks [get]
func GetAllTask(c *gin.Context) {
	logrus.Info("GetAllTask endpoint")
	var tasks []models.Task
	resultChan := make(chan []models.Task)
	pool.wg.Add(1)
	go func() {
		defer pool.wg.Done()
		initializers.DB.Find(&tasks)
		resultChan <- tasks
	}()

	tasks = <-resultChan
	logrus.Infof("Tasks are listed successfully.")
	c.JSON(200, gin.H{"tasks": tasks})
}

// TaskShowByID returns a JSON response with the task specified by ID.
// @Summary Get a task by ID
// @Tags Tasks
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} models.Task
// @Router /tasks/{id} [get]
func TaskShowByID(c *gin.Context) {
	logrus.Info("TaskShowByID endpoint")

	id := c.Param("id")
	// Get Task by ID
	var task models.Task
	if err := initializers.DB.First(&task, id).Error; err != nil {
		logrus.WithError(err).Error("Failed to process task show by id (task not found)")
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	// Response
	logrus.Infof("Task listed successfully.")
	c.JSON(200, gin.H{"task": task})
}

// TaskUpdate updates a task specified by ID in the database.
// It returns a JSON response indicating the update.
// @Summary Update a task by ID
// @Tags Tasks
// @Produce json
// @Param id path int true "Task ID"
// @Param title formData string true "Updated task title"
// @Param description formData string true "Updated task description"
// @Param status formData string true "Updated task status"
// @Success 200 {object} models.Task
// @Router /tasks/{id} [put]
func TaskUpdate(c *gin.Context) {
	logrus.Info("TaskUpdate endpoint")
	// Get ID
	id := c.Param("id")

	// Get existing task by ID
	var existingTask models.Task
	if err := initializers.DB.First(&existingTask, id).Error; err != nil {
		logrus.WithField("id", id).WithError(err).Error("Failed to get existing task for updating")
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	// Bind the updated data from the request form
	var updatedData struct {
		Title       string `form:"title" binding:"required"`
		Description string `form:"description" binding:"required"`
		Status      string `form:"status" binding:"required"`
	}
	if err := c.ShouldBind(&updatedData); err != nil {
		logrus.WithError(err).Error("Failed to bind request body")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Update task fields
	existingTask.Title = updatedData.Title
	existingTask.Description = updatedData.Description
	existingTask.Status = updatedData.Status

	// Save the updated task to the database
	if err := initializers.DB.Save(&existingTask).Error; err != nil {
		logrus.WithError(err).Error("Failed to update task")
		c.JSON(500, gin.H{"error": "Failed to update task"})
		return
	}

	// Response with the updated task
	logrus.Infof("Task updated successfully.")
	c.JSON(200, gin.H{"task": existingTask})
}

// TaskDelete deletes a task specified by ID from the database.
// It returns a JSON response indicating the deletion.
// @Summary Delete a task by ID
// @Tags Tasks
// @Param id path int true "Task ID"
// @Success 200
// @Router /tasks/{id} [delete]
func TaskDelete(c *gin.Context) {
	logrus.Info("TaskDelete endpoint")
	// Get ID
	id := c.Param("id")

	// Delete Task by ID
	if err := initializers.DB.Delete(&models.Task{}, id).Error; err != nil {
		logrus.WithError(err).Error("Failed to delete task")
		c.JSON(500, gin.H{"error": "Failed to delete task"})
		return
	}
	// Response
	logrus.Infof("Task deleted successfully.")
	c.Status(200)
}
