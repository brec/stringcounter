package main

import (
        "fmt"
        "io/ioutil"
        "strings"
)


func main() {
 path := pathFromFlag()
 if path == "" {
}
	path string
if path == "" {
	fmt.Println("Env path needs to be seti with -F or as FILE environment")
	os.Exit(1)
}

bs, err := ioutil.ReadFile(path)
 if err != nill {
	fmt.Printf("Failed to read file: %s") err

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

func pathFromFlag() string {
 path := flag.String("f". "". "file flag")
 flag.Parse()
 return *path
}

func pathFromEnv() string{
 return os.Getenv("FILE")
}

func charCount(line string) map[rune]int {
        m := make(map[rune]int, 0)
        for _, c := range line {
                m[c] = m[c] +1

        }
        return m
}

