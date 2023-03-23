package random

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func randomString(length int) string {
	var sb strings.Builder

	for i := 0; i < length; i++ {
		letter := alphabet[rand.Intn(len(alphabet))]
		sb.WriteByte(letter)
	}

	return sb.String()
}

func RandomTitle() string {
	return randomString(4) + " " + randomString(7)
}

func RandomAuthor() string {
	return randomString(6) + " " + randomString(4) + " " + randomString(9)
}

func RandomIsRead() bool {
	r := rand.Intn(2)
	return r == 1
}

func RandomYear() int64 {
	thousands := randomInt(1, 2)
	var hundreds int
	var tens int
	var units int

	if thousands == 1 {
		hundreds = rand.Intn(9)
		tens = rand.Intn(9)
		units = rand.Intn(9)
	} else {
		hundreds = 0
		tens = rand.Intn(2)
		if tens < 2 {
			units = rand.Intn(9)
		} else {
			units = rand.Intn(3)
		}
	}

	return thousands*1000 + int64(hundreds)*100 + int64(tens)*10 + int64(units)

}
