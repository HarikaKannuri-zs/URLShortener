package utils

import (
	"math/rand"
	"time"
)

const CharSet = "abcdefghijklmnopqrstuvwxyz"
const Aliaslength = 6

func init() {
	rand.New(rand.NewSource((time.Now().UnixNano())))
}

func GenerateRandomAlias() string {
	b := make([]byte, Aliaslength)
	for i := range b {
		b[i] = CharSet[rand.Intn(len(CharSet))]
	}
	return string(b)
}
