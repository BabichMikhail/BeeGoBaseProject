package utils

import (
	"math/rand"
	"time"
)

func Randomize(additionalToNowTime ...int64) {
	if len(additionalToNowTime) == 1 {
		rand.Seed(time.Now().UTC().UnixNano() + additionalToNowTime[0])
	} else {
		rand.Seed(time.Now().UTC().UnixNano())
	}
}
