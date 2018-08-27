package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)


func main() {
	filename := os.Args[1:] 
	data, err := ioutil.ReadFile(filename)
		if err != nil{
			fmt.Fprintf(os.Stderr, "stringcounter: %v\n", err)
			continue
		}
		lines := strings.Split(filename, "\n")
		// /n character added because print doesnt add it by default
		fmt.Printf("%s\n", l)
		for k, v := range charCount(l) {
			fmt.Printf("'%c'=%d, ", k, v)
		}
		fmt.Print("\n\n")
	}

func charCount(line string) map[rune]int {
	m := make(map[rune]int, 0)
	for _, c := range line {
		m[c] = m[c] +1

	}
	return m
}


