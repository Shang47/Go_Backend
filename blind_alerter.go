package poker

import (
	"fmt"
	"io"
	"time"
)

// BlindAlerter schedules alerts for blind amounts.
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int, to io.Writer)
}

// BlindAlerterFunc allows you to implement BlindAlerter with a function.
type BlindAlerterFunc func(duration time.Duration, amount int, to io.Writer)

// ScheduleAlertAt is BlindAlerterFunc's implementation of BlindAlerter.
func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	a(duration, amount, to)
}

// StdOutAlerter returns a BlindAlerterFunc that schedules alerts and prints them to out.
func Alerter(duration time.Duration, amount int, to io.Writer) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(to, "Blind is now %d\n", amount)
	})
}
func NewAlerter() (BlindAlerterFunc, <-chan string) {
	alerts := make(chan string)

	scheduleAlertAt := func(duration time.Duration, amount int, to io.Writer) {
		time.AfterFunc(duration, func() {
			alerts <- fmt.Sprintf("Blind is now %d", amount)
		})
	}
	return scheduleAlertAt, alerts
}
