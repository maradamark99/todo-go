package scheduler

import (
	"time"
)

type ToSchedule func()

func ScheduleOnce(duration time.Duration, fn ToSchedule) {
	timer := time.NewTimer(duration)
	<-timer.C
	fn()
}
