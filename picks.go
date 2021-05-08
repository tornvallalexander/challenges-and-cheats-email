package main

import (
	"math/rand"
	"time"
)

func GetRandomPicks(array []string) []string {
	picks := make([]string, 0)
	// to make sure we don't get stuck in a never-ending loop for whatever reason
	bigNumber := 1000

	for i := 0; i < bigNumber; i++ {
		randomIndex := GetRandomNumber(len(array))
		picks = append(picks, array[randomIndex])

		if len(picks) == 2 && picks[0] != picks[1] {
			break
		}

		// ensures different values
		if len(picks) == 2 {
			picks = picks[:len(picks) - 1]
		}
	}

	return picks
}

// GetRandomNumber returns a completely random number, a non-deterministic one
func GetRandomNumber(length int) int {
	// new seed
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	randIndex := random.Intn(length - 1)

	return randIndex
}
