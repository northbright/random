package random

import (
	"fmt"
	"math/rand"
	"time"
)

// RandTime returns random time between minimize time + pad and maximum time - pad.
func RandTime(minTime, maxTime time.Time, pad time.Duration) (time.Time, error) {
	if minTime.After(maxTime) {
		return time.Now(), fmt.Errorf("min time > max time")
	}

	min := minTime.Add(pad)
	max := maxTime.Add(0 - pad)

	if min.After(max) {
		return time.Now(), fmt.Errorf("min time + pad > max time - pad")
	}

	d := max.Sub(min)

	nMillisecond := int64(d / time.Millisecond)
	if nMillisecond <= 0 {
		return time.Now(), nil
	}

	rand.Seed(time.Now().UnixNano())
	x := rand.Int63n(nMillisecond)

	return min.Add(time.Duration(x) * time.Millisecond), nil
}
