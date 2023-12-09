package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/musaubrian/aoc-23/utils"
)

func main() {
	var largerArr [][]string
	var ttl int
	vals := utils.ReadFile("input.txt")

	for _, v := range vals {
		var nums []string
		for _, r := range v {
			if unicode.IsDigit(r) {
				nums = append(nums, string(r))
			}
		}
		largerArr = append(largerArr, nums)
	}
	for _, q := range largerArr {
		if len(q) > 1 {
			n := q[0] + q[len(q)-1]
			res, _ := strconv.Atoi(n)
			ttl += res

		} else {
			n := q[0] + q[0]
			res, _ := strconv.Atoi(n)
			ttl += res
		}
	}
	fmt.Println(ttl)
}
