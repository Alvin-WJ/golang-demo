package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	/* Word regular experssion. */
	re, _ := regexp.Compile("[a-zA-Z0-9]+")
	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err != os.EOF {
				fmt.Fprintf(os.Stderr, "error: can't read - %s\n", err.String())
			}
			break
		}
		matches := re.FindAll(line, -1)
		for _, word := range(matches) {
			fmt.Printf("%s\t1\n", word)
		}
	}
}