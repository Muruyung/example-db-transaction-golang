package utils

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomTransactionCode() string {
	dt := time.Now()
	min := math.Pow(10, 4-1)
	max := math.Pow(10, 4) - 1

	randInteger := strconv.Itoa(randInt(int(min), int(max)))
	result := dt.Format("200602011504") + randInteger

	return result
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}
