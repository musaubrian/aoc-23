package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/musaubrian/aoc-23/utils"
)

func main() {
	vals := utils.ReadFile("input")
	result := 0
	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	for gameId, gameSets := range vals {
		gamePossible := true
		parsedGamesSets := strings.Split(strings.TrimPrefix(gameSets, fmt.Sprintf("Game %d: ", gameId+1)), "; ")

		for _, gameSet := range parsedGamesSets {
			for _, gameSetDice := range strings.Split(gameSet, ", ") {
				digit, err := strconv.ParseInt(strings.Split(gameSetDice, " ")[0], 10, 0)
				if err != nil {
					log.Fatal(err)
				}
				if strings.HasSuffix(gameSetDice, "red") {
					if maxRed < int(digit) {
						gamePossible = false
						break
					}
				} else if strings.HasSuffix(gameSetDice, "green") {
					if maxGreen < int(digit) {
						gamePossible = false
						break
					}
				} else if strings.HasSuffix(gameSetDice, "blue") {
					if maxBlue < int(digit) {
						gamePossible = false
						break
					}
				} else {
					panic(1)
				}
			}
		}
		if gamePossible {
			result = result + (gameId + 1)
		}
	}

	println(result)
}
