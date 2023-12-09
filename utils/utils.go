package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(path string) []string {
	var vals []string
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		vals = append(vals, sc.Text())
	}
	return vals
}
