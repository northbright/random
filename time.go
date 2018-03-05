package random

import (
	"fmt"
	"math/rand"
	"time"
)

// Time returns random time between minimize time and maximum time.
func Time(minTime, maxTime time.Time) (time.Time, error) {
	if minTime.After(maxTime) {
		return time.Now(), fmt.Errorf("min time > max time")
	}

	d := maxTime.Sub(minTime)

	nMillisecond := int64(d / time.Millisecond)
	if nMillisecond <= 0 {
		return minTime, nil
	}

	rand.Seed(time.Now().UnixNano())
	x := rand.Int63n(nMillisecond)

	return minTime.Add(time.Duration(x) * time.Millisecond), nil
}
