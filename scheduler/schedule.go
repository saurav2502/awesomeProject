package scheduler

import (
	"fmt"
	"github.com/mileusna/crontab"
	"log"
)

func Init() {
	ctab := crontab.New()
	err := ctab.AddJob("*/5 * * * *", myFunc) // on 1st day of month
	if err != nil {
		log.Println(err)
		return
	}
}

func myFunc() {
	fmt.Println("Helo, world")
}
