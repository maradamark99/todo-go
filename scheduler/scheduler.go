package scheduler

import (
	"time"
)

type ToSchedule func()

func Schedule(duration time.Duration, fn ToSchedule) {
	timer := time.NewTimer(duration)
	<-timer.C
	fn()
}
