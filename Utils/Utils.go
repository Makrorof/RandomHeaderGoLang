package Utils

import (
	"math/rand"
	"time"
)

func Init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int) int {
	// set seed
	//rand.Seed(time.Now().UnixNano())
	// generate random number and print on console
	return rand.Intn(max-min) + min
}
