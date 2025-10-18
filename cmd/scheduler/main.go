package main

import (
	"time"

	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"github.com/TroJanBoi/assembly-visual-backend/internal/scheduler"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithLocation(time.FixedZone("Asia/Bangkok", 7*3600)))
	c.AddFunc("0 3 * * *", func() {
		scheduler.CleanupSoftDeletedUsers(&model.User{}, 7)
		// scheduler.CleanupSoftDeletedRecords(&model.Room{}, 30)
		scheduler.CleanupExpiredInvitations(&model.Invitation{}, 1)
	})
	c.Start()
	select {}
}
