package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := ioutil.ReadFile("file.txt")
	check(err)

	goProverbs := string(f) // f is a byte slice

	lines := strings.Split(goProverbs, "\n")

	for _, line := range lines {
		fmt.Println(line)

		for k, v := range charCount(line) {
			fmt.Printf("'%s'=%d, ", k, v)
		}

		fmt.Print("\n\n")
	}
}

func charCount(line string) map[string]int {
	m := make(map[string]int, 0)
	for _, char := range line {
		m[string(char)] = m[string(char)] + 1
	}

	return m
}
