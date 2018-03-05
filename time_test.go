package random_test

import (
	"log"
	"time"

	"github.com/northbright/random"
)

func ExampleTime() {
	testDatas := []struct {
		Min time.Time
		Max time.Time
	}{
		{
			time.Now(),
			time.Now(),
		},
		{
			time.Now(),
			time.Now().Add(time.Second),
		},
		{
			time.Now(),
			time.Now().Add(time.Second * 1000),
		},
	}

	for i, testData := range testDatas {
		log.Printf("\n-----------case %v:", i)
		t, err := random.Time(testData.Min, testData.Max)
		if err != nil {
			log.Printf("Time() error: %v\nt: %v\nmin: %v\nmax: %v\n", err, t, testData.Min, testData.Max)
		} else {
			log.Printf("Time() OK\nt: %v\nmin: %v\nmax: %v\n", t, testData.Min, testData.Max)
		}
	}
	// Output:
}
