package scheduler

import (
	"log"
	"time"

	"github.com/TroJanBoi/assembly-visual-backend/internal/database"
)

func CleanupSoftDeletedUsers(model any, days int) {
	db := database.New().GetClient()
	// threshold := time.Now().AddDate(0, 0, -days)
	threshold := time.Now().Add(-1 * time.Minute)

	var count int64
	if err := db.Unscoped().Model(model).Where("deleted_at IS NOT NULL AND deleted_at < ?", threshold).Count(&count).Error; err != nil {
		log.Printf("❌ Failed to count old records for %+v: %v", model, err)
		return
	}

	if count == 0 {
		log.Printf("ℹ️ No old records to cleanup for %+v", model)
		return
	}
	if err := db.Unscoped().Where("deleted_at IS NOT NULL AND deleted_at < ?", threshold).Delete(model).Error; err != nil {
		log.Printf("❌ Failed to cleanup old records for %+v: %v", model, err)
		return
	}
	log.Printf("✅ Successfully cleaned up %d old records for %+v", count, model)
}
