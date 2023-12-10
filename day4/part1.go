package main

import (
	"fmt"
	"strings"

	"github.com/musaubrian/aoc-23/utils"
)

// Help me visualize stuff easier
type Card struct {
	available []string
	winning   []string
}

func main() {
	vals := utils.ReadFile("input")
	// vals := utils.ReadFile("d4")
	var cards []Card
	var count int
	for _, v := range vals {
		sep := strings.Split(v, ":")
		nums := strings.Split(sep[1], "|")

		cards = append(cards, Card{
			available: strings.Split(nums[1], " "),
			winning:   strings.Split(nums[0], " "),
		})
	}

	for _, card := range cards {
		var present int
		for _, v := range card.winning {
			for _, y := range card.available {
				if v == y {
					if v != "" || y != "" {
						present++
					}
				}
			}
		}
		if present == 2 {
			present = 2
			count += present
		} else if present == 1 {
			present = 1
			count += present
		} else if present > 2 {
			present := 1 << (present - 1)
			count += present
		}
	}
	fmt.Println("COUNT: ", count)
}
