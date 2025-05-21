package cron

import (
	"log"

	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/controller"
	"github.com/robfig/cron/v3"
)

func Init() {
	c := cron.New(cron.WithSeconds())

	// Controllers
	cronRecoveryCode := controller.NewCronRecoveryCode()
	cronRecoveryToken := controller.NewCronRecoveryToken()
	cronAuth := controller.NewCronAuth()

	_, err := c.AddFunc("@every 5m", cronRecoveryCode.CheckRecoveryCodes)
	if err != nil {
		log.Fatalf("Failed to schedule cronjob: %v", err)
	}
	_, err = c.AddFunc("@every 5m", cronRecoveryToken.CheckRecoveryTokens)
	if err != nil {
		log.Fatalf("Failed to schedule cronjob: %v", err)
	}
	_, err = c.AddFunc("@daily", cronAuth.DeleteExpiredSessions)
	if err != nil {
		log.Fatalf("Failed to schedule cronjob: %v", err)
	}
	_, err = c.AddFunc("@daily", cronAuth.DeleteTokens)
	if err != nil {
		log.Fatalf("Failed to schedule cronjob: %v", err)
	}

	c.Start()

	log.Println("Cronjob started.")

	select {}
}
