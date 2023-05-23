package initializers

import (
	"fmt"

	"github.com/imnmania/go_crud/models"
)

func SyncDatabase() {
	err := DB.AutoMigrate(&models.Post{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Database migration complete...")
	}
}
