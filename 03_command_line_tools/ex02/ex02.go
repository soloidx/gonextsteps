package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type proverb struct {
	line  string
	chars map[rune]int
}

func (p *proverb) countChars() {
	if p.chars != nil {
		return
	}

	p.chars = make(map[rune]int, 0)

	for _, char := range p.line {
		p.chars[char] = p.chars[char] + 1
	}
}

func newProverb(line string) *proverb {
	p := new(proverb)
	p.line = line
	p.countChars()
	return p
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func main() {
	path := pathFromFlag()

	if path == "" {
		fmt.Println("You must specify one file with the -f")
		os.Exit(1)
	}

	proverbs, err := loadProverbs(path)
	if err != nil {
		fmt.Printf("Failed to load proverbs: %s", err)
		os.Exit(1)
	}

	for _, p := range proverbs {
		fmt.Println(p.line)

		for k, v := range p.chars {
			fmt.Printf("'%c'=%d, ", k, v)
		}

		fmt.Print("\n\n")
	}
}

func pathFromFlag() string {
	path := flag.String("f", "", "text file path") //return a *string
	flag.Parse()
	return *path
}

func loadProverbs(path string) ([]*proverb, error) {
	var proverbs []*proverb

	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(f), "\n")

	for _, line := range lines {
		proverbs = append(proverbs, newProverb(line))
	}

	return proverbs, nil
}
