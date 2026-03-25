package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/TroJanBoi/assembly-visual-backend/internal/database"
	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"gorm.io/datatypes"
)

// =========================================================================
// CONFIGURATION
// Replace MY_EMAIL with the actual Gmail address you use to log in.
// This ensures that when you log into the platform, you will see 
// the generated data tied to your Google OAuth account.
// =========================================================================
const MY_EMAIL = "netsdev3@gmail.com"
const MY_NAME = "Prakan Suma"

func main() {
	rand.Seed(time.Now().UnixNano())
	log.Println("Starting mock data generation for Assembly Simulator Platform...")

	// Initialize database
	dbService := database.New()
	if dbService == nil {
		log.Fatal("Failed to initialize database")
	}
	db := dbService.GetClient()
	defer dbService.Close()

	// Hard delete old mock info
	db.Unscoped().Where("email LIKE ?", "sim_user_%@test.com").Delete(&model.User{})
	db.Unscoped().Where("email LIKE ?", "mock_user_%@test.com").Delete(&model.User{})

	log.Println("Creating 19 mock students and 1 specific user...")
	users := make([]model.User, 20)

	// User 0 is YOUR specific user
	users[0] = model.User{
		Email:        MY_EMAIL,
		Name:         MY_NAME,
		PasswordHash: "mocked_hash",
	}

	// Users 1-19 are mock students
	for i := 1; i < 20; i++ {
		users[i] = model.User{
			Email:        fmt.Sprintf("sim_user_%d@test.com", i),
			Name:         fmt.Sprintf("Assembly Student %d", i),
			PasswordHash: "mocked_hash",
		}
	}

	// We'll also manually create a Teacher user if your account isn't going to be the owner of ALL classes
	teacher := model.User{
		Email:        "sim_teacher@test.com",
		Name:         "Prof. Von Neumann",
		PasswordHash: "mocked_hash",
	}

	// Safely create or find users to avoid unique constraint violations
	for i := range users {
		db.Where("email = ?", users[i].Email).FirstOrCreate(&users[i])
	}
	db.Where("email = ?", teacher.Email).FirstOrCreate(&teacher)
	
	allUsers := append(users, teacher)
	
	// Refresh the yourUser variable to have the database ID
	var yourUser model.User
	db.Where("email = ?", MY_EMAIL).First(&yourUser)
	var profUser model.User
	db.Where("email = ?", "sim_teacher@test.com").First(&profUser)

	log.Println("Creating 10 Classrooms...")
	classNames := []string{
		"Computer Architecture 101",
		"Advanced Assembly",
		"Microprocessors Laboratory",
		"8-bit System Design",
		"Low-Level Programming",
		"Embedded Systems Intro",
		"Digital Logic & Assembly",
		"Reverse Engineering basics",
		"Compiler Construction",
		"Operating Systems 101",
	}

	var classrooms []model.Classroom
	for i, name := range classNames {
		// Make YOU the owner of half the classes, and the Professor the owner of the other half
		var ownerID int
		if i%2 == 0 {
			ownerID = int(yourUser.ID)
		} else {
			ownerID = int(profUser.ID)
		}

		classrooms = append(classrooms, model.Classroom{
			OwnerId:     ownerID,
			Topic:       name,
			Description: fmt.Sprintf("Learn about %s using our 8-bit simulator.", name),
			Status:      0,
			Code:        fmt.Sprintf("CLASS-%02d", i+1),
		})
	}
	if err := db.Create(&classrooms).Error; err != nil {
		log.Fatalf("Failed to create classrooms: %v", err)
	}

	log.Println("Adding members to classrooms (3-13 students per class)...")
	var members []model.Member
	var yourClassIds []int
	
	for _, c := range classrooms {
		// Ensure the owner is a teacher in their own class
		members = append(members, model.Member{
			UserID:  c.OwnerId,
			ClassID: int(c.ID),
			Role:    "teacher",
			JoinAt:  time.Now(),
		})

		// You should ideally be enrolled in EVERY class so you can see them all in screenshots
		if c.OwnerId != int(yourUser.ID) {
			members = append(members, model.Member{
				UserID:  int(yourUser.ID),
				ClassID: int(c.ID),
				Role:    "member",
				JoinAt:  time.Now(),
			})
			yourClassIds = append(yourClassIds, int(c.ID))
		} else {
			yourClassIds = append(yourClassIds, int(c.ID))
		}

		// Randomly add 3 to 13 mock members
		numMembers := rand.Intn(11) + 3 // 3 to 13
		
		// Shuffle mock users
		mockIndices := rand.Perm(19)
		
		for i := 0; i < numMembers; i++ {
			studentIndex := mockIndices[i] + 1 // offset to 1-19
			members = append(members, model.Member{
				UserID:  int(allUsers[studentIndex].ID),
				ClassID: int(c.ID),
				Role:    "member",
				JoinAt:  time.Now(),
			})
		}
	}
	if err := db.Create(&members).Error; err != nil {
		log.Fatalf("Failed to create members: %v", err)
	}

	log.Println("Creating assignments (2-5 per class)...")
	var assignments []model.Assignment
	for _, c := range classrooms {
		numAssignments := rand.Intn(4) + 2 // 2 to 5
		for j := 0; j < numAssignments; j++ {
			assignments = append(assignments, model.Assignment{
				ClassID:     int(c.ID),
				Title:       fmt.Sprintf("Assignment %d: CPU Tasks", j+1),
				Description: "Load the correct values into the registers and halt the CPU.",
				Grade:       10,
				MaxAttempt:  5,
				Setting:     datatypes.JSON(`{"cpu_clock_speed": 100}`),
				Condition:   datatypes.JSON(fmt.Sprintf(`{"type": "register_eq", "register": "A", "value": %d}`, j*5)),
			})
		}
	}
	if err := db.Create(&assignments).Error; err != nil {
		log.Fatalf("Failed to create assignments: %v", err)
	}

	log.Println("Creating playgrounds and submissions...")
	var playgrounds []model.Playground
	var submissions []model.Submission

	for _, a := range assignments {
		// Find members for this class
		var classMembers []model.Member
		for _, m := range members {
			if m.ClassID == a.ClassID && m.Role == "member" {
				classMembers = append(classMembers, m)
			}
		}

		for _, m := range classMembers {
			// Assembly syntax example
			assemblyCode := fmt.Sprintf("MOV A, %d\\nADD A, 5\\nOUT 1, A\\nHLT", rand.Intn(10))

			pg := model.Playground{
				AssignmentID: int(a.ID),
				UserID:       m.UserID,
				Item:         datatypes.JSON(fmt.Sprintf(`{"code": "%s"}`, assemblyCode)),
				Status:       "completed",
			}
			playgrounds = append(playgrounds, pg)
		}
	}

	if len(playgrounds) > 0 {
		if err := db.Create(&playgrounds).Error; err != nil {
			log.Fatalf("Failed to create playgrounds: %v", err)
		}
	}

	// Create submissions for the playgrounds
	for _, pg := range playgrounds {
		pass := rand.Intn(10) > 2 // 70% pass rate
		score := 0.0
		status := "failed"
		clientRes := `{"success": false}`
		serverRes := `{"passed": false}`
		if pass {
			score = 10.0
			status = "verified"
			clientRes = `{"success": true}`
			serverRes = `{"passed": true, "score": 10}`
		}

		sub := model.Submission{
			UserID:        pg.UserID,
			AssignmentID:  pg.AssignmentID,
			PlaygroundID:  int(pg.ID),
			AttemptNumber: 1,
			ItemSnapshot:  pg.Item,
			ClientResult:  datatypes.JSON(clientRes),
			ServerResult:  datatypes.JSON(serverRes),
			Score:         score,
			Status:        status,
			IsVerified:    pass,
			DurationMS:    rand.Intn(200) + 20,
			FeedBack:      "Auto-graded by 8-bit simulator.",
		}
		submissions = append(submissions, sub)
	}
	if len(submissions) > 0 {
		if err := db.Create(&submissions).Error; err != nil {
			log.Fatalf("Failed to create submissions: %v", err)
		}
	}

	log.Println("Platform-specific mock data generation completed successfully!")
}
