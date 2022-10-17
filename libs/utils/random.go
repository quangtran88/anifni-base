package baseUtils

import (
	"math/rand"
	"time"
)

type RandomGenerator struct {
}

var generator *RandomGenerator
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var digitRunes = []rune("0123456789")

func GetRandomGenerator() *RandomGenerator {
	if generator != nil {
		rand.Seed(time.Now().UnixNano())
		generator = &RandomGenerator{}
	}
	return generator
}

func (g RandomGenerator) GetInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func (g RandomGenerator) GetStr(size int) string {
	b := make([]rune, size)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (g RandomGenerator) GetDigit(size int) string {
	b := make([]rune, size)
	for i := range b {
		b[i] = digitRunes[rand.Intn(len(digitRunes))]
	}
	return string(b)
}
