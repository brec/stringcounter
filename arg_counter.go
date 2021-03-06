package main

import (
        "fmt"
        "io/ioutil"
        "strings"
)


func main() {
	bs, err := ioutil.ReadFile("proverbs.txt")
	if err != nil {
		panic(fmt.Errorf("failed to read file: %s", err))
}
proverbs := string(bs)
                
	lines := strings.Split(proverbs, "\n")
         // /n character added because print doesnt add it by default
	for _, l := range lines {
		fmt.Printf("%s\n", l)
                for k, v := range charCount(l) {
                        fmt.Printf("'%c'=%d, ", k, v)
                }
                fmt.Print("\n\n")
        }
}

func charCount(line string) map[rune]int {
        m := make(map[rune]int, 0)
        for _, c := range line {
                m[c] = m[c] +1

        }
        return m
}

