package main

import (
	"context"
	"fmt"
	"log"

	"github.com/TroJanBoi/assembly-visual-backend/internal/database"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
)

func main() {
	dbService := database.New()
	if dbService == nil {
		log.Fatal("Failed to DB")
	}
	db := dbService.GetClient()
	defer dbService.Close()

	repo := repository.NewUserRepository(db)

	classes, err := repo.GetMeClass(context.Background(), 3)
	if err != nil {
		log.Fatalf("Error GetMeClass: %v", err)
	}

	fmt.Printf("GetMeClass success! Found %d classes\n", len(*classes))
}
