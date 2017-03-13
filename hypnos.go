package hypnos

import (
	"time"

	"github.com/gorhill/cronexpr"
)

// Sleep returns the length of time to wait until the next interval expressed in
// crontab notation along with the date for when that will happen. If the crontab
// expression fails to parse an error will be returned.
func Sleep(interval string) (time.Duration, time.Time, error) {
	var duration time.Duration
	var date time.Time
	expr, err := cronexpr.Parse(interval)
	if err != nil {
		return duration, date, err
	}
	date = expr.Next(time.Now())
	duration = date.Sub(time.Now())
	time.Sleep(duration)
	return duration, date, nil
}
