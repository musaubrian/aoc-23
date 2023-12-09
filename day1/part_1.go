package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	var vals []string
	var largerArr [][]string
	var ttl int

	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		vals = append(vals, sc.Text())
	}
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
