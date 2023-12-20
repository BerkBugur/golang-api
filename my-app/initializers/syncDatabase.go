package initializers

import (
	"github.com/BerkBugur/Go-Project/models"
)

func SyncDatabase() {

	DB.AutoMigrate(&models.Task{})
	DB.AutoMigrate(&models.Users{})
}
