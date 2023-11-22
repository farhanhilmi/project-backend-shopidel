package util

import (
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron"
)

func RunCronJobs() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(6).Hour().Do(func() {
		err := os.RemoveAll("./imageuploads/")
		if err != nil {
			log.Println("Error remove imageuploads:", err)
		}
	})

	s.StartAsync()
}
