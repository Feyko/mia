package mia

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"log"
	"os"
	"os/exec"
)

// Watches the media resources, effectively running this program at the specified interval, at 5pm
func Watch(days int) {
	ScheduleCron(fmt.Sprintf("0 17 */%v * *", days))
}

// Schedule mia to run directly with a cron schedule string
func ScheduleCron(schedule string) {
	home, _ := homedir.Dir()
	cronfilePath := home + "/.mia.cron"

	f, err := os.OpenFile(cronfilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		log.Fatalf("Failed to open cron file: %v", err)
	}
	defer f.Close()

	_, err = f.WriteString(schedule + " mia\n")
	if err != nil {
		log.Fatalf("Failed to write to cron file: %v", err)
	}

	o, err := exec.Command(cronfilePath).Output()
	if err != nil {
		fmt.Println(o)
		log.Fatalf("Error scheduling cron: %v", err)
	}
}
