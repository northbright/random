package random_test

import (
	"log"
	"time"

	"github.com/northbright/random"
)

func ExampleRandTime() {
	testDatas := []struct {
		Min time.Time
		Max time.Time
		Pad time.Duration
	}{
		{
			time.Now(),
			time.Now(),
			0,
		},
		{
			time.Now(),
			time.Now(),
			time.Millisecond,
		},
		{
			time.Now().Add(time.Second),
			time.Now(),
			0,
		},
		{
			time.Now(),
			time.Now().Add(time.Second),
			time.Millisecond * 600,
		},
		{
			time.Now(),
			time.Now().Add(time.Second),
			time.Millisecond * 500,
		},
		{
			time.Now(),
			time.Now().Add(time.Second),
			time.Millisecond * 200,
		},
		{
			time.Now(),
			time.Now().Add(time.Second * 1000),
			time.Millisecond * 200,
		},
	}

	for i, testData := range testDatas {
		log.Printf("\n-----------case %v:", i)
		t, err := random.RandTime(testData.Min, testData.Max, testData.Pad)
		if err != nil {
			log.Printf("RandTime() error: %v\nt: %v\nmin: %v\nmax: %v\npad: %v\n", err, t, testData.Min, testData.Max, testData.Pad)
		} else {
			log.Printf("RandTime() OK\nt: %v\nmin: %v\nmax: %v\npad: %v\n", t, testData.Min, testData.Max, testData.Pad)
		}
	}
	// Output:
}
