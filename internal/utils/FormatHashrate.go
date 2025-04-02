package utils

import (
	"fmt"
	"time"
)

func FormatHashrate(numTries uint64, start time.Time) string {
	t := time.Now()
	elapsed := t.Sub(start)
	hasrate := float64(numTries) * float64(time.Second) / float64(elapsed)
	return fmt.Sprintf("%f H/s", hasrate)

}
