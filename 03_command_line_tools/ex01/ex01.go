package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You must specify one argument as the file path.")
		os.Exit(1)
	}

	path := os.Args[1]

	f, err := ioutil.ReadFile(path)
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
