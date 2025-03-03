package local_time

import (
	"log"
	"time"
)

var Location *time.Location

func InitTime() {
	var err error
	Location, err = time.LoadLocation("Asia/Dhaka")
	if err != nil {
		log.Fatal(err)
	}
}
