package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	var src, change string
	var n, start, end int
	fmt.Fscan(input, &src)
	fmt.Fscan(input, &n)
	byf := []byte(src)
	for n > 0 {
		fmt.Fscan(input, &start)
		fmt.Fscan(input, &end)
		fmt.Fscan(input, &change)
		start--
		for i := range change {
			byf[start] = change[i]
			start++
		}
		n--
	}
	output.WriteString(string(byf))
	output.WriteString("\n")
}
