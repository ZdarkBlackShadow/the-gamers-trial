package config

import (
	"math/rand"
	"time"
)

var Rand *rand.Rand

func InitRandomGenerator() {
	source := rand.NewSource(time.Now().UnixNano())
	Rand = rand.New(source)
}
