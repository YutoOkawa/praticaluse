package duration

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	tz, _ := time.LoadLocation("America/Los_Angeles")
	future := time.Date(2015, time.October, 21, 7, 28, 0, 0, tz)
	fmt.Println(now.String())
	fmt.Println(future.String())

	fiveMinute := 5 * time.Minute
	var seconds int = 10
	tenSeconds := time.Duration(seconds) * time.Second

	past := time.Date(1995, time.November, 112, 6, 38, 0, 0, time.UTC)
	dur := time.Now().Sub(past)

	fiveMinuteAfter := time.Now().Add(fiveMinute)
	fiveMinuteBefore := time.Now().Add(-fiveMinute)
}
