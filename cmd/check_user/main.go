package main

import (
	"fmt"
	"log"

	"github.com/TroJanBoi/assembly-visual-backend/internal/database"
	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
)

func main() {
	dbService := database.New()
	if dbService == nil {
		log.Fatal("Failed to db")
	}
	db := dbService.GetClient()
	defer dbService.Close()

	// Mimic GetMeClass precisely
	var member []model.Member
	if err := db.Where("user_id = ?", 3).Find(&member).Error; err != nil {
		log.Fatal(err)
	}

	var classIDs []int
	for _, m := range member {
		classIDs = append(classIDs, int(m.ClassID))
	}

	var classes []model.Classroom
	if err := db.Where("id IN ?", classIDs).Find(&classes).Error; err != nil {
		log.Fatal(err)
	}

	for _, class := range classes {
		var user model.User
		if err := db.Where("id = ?", class.OwnerId).First(&user).Error; err != nil {
			log.Fatalf("FAILED retrieving owner for class ID %d (Topic: %s, OwnerId expected: %d): %v", class.ID, class.Topic, class.OwnerId, err)
		}
		fmt.Printf("Sucess for class %d, Owner is %s\n", class.ID, user.Name)
	}
}
