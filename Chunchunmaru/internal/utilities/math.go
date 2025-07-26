package utilities

import (
	"math/rand"
	"time"
)

func RandomNumber(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandomDuration(min, max time.Duration) time.Duration {
	if max < min {
		min, max = max, min // swap to ensure min <= max
	}
	delta := max - min
	if delta == 0 {
		return min
	}
	// rand.Int63n returns [0, delta)
	return min + time.Duration(rand.Int63n(int64(delta)+1))
}
