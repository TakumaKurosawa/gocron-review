package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func main() {
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Println(err)

		return
	}

	j, err := s.NewJob(
		gocron.CronJob("* * * * * *", true),
		gocron.NewTask(
			func(str string) {
				if err := say(str); err != nil {
					log.Println(err)
				}
			},
			"Hello, World!",
		),
	)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(j.ID())

	s.Start()

	select {
	case <-time.After(time.Minute):
	}

	if err := s.Shutdown(); err != nil {
		log.Println(err)
	}
}

func say(str string) error {
	fmt.Println(str)

	return nil
}
