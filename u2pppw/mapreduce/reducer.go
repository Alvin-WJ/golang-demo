package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	counts := make(map[string]uint)
	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err != os.EOF {
				fmt.Fprintf(os.Stderr, "error: can't read - %s\n", err)
			}
			break
		}
		i := bytes.IndexByte(line, '\t')
		if i == -1 {
			fmt.Fprintln(os.Stderr, "error: can't find tab")
			continue
		}
		word := string(line[0:i])
		count, err := strconv.Atoui(string(line[i+1:]))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error: bad number - %s\n", err)
			continue
		}

		counts[word] = counts[word] + count
	}

	/* Output aggregated counts. */
	for word, count := range(counts) {
		fmt.Printf("%s\t%d\n", word, count)
	}
}
