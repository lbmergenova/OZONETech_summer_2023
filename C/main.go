package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	var t, n, temp int
	var str string
	fmt.Fscan(input, &t)
	for t > 0 {
		interval := [2]int{15, 30}
		fmt.Fscan(input, &n)
		for n > 0 {
			fmt.Fscan(input, &str)
			fmt.Fscan(input, &temp)
			if interval[0] != 0 {
				if interval[0] != 0 && str == "<=" && temp >= interval[0] {
					if temp < interval[1] {
						interval[1] = temp
					}
					output.WriteString(strconv.Itoa(interval[1]))
				} else if interval[0] != 0 && str == ">=" && temp <= interval[1] {
					if temp > interval[0] {
						interval[0] = temp
					}
					output.WriteString(strconv.Itoa(interval[1]))
				} else {
					interval[0] = 0
					output.WriteString("-1")
				}
			} else {
				output.WriteString("-1")
			}
			output.WriteString("\n")
			n--
		}
		output.WriteString("\n")
		t--
	}
}
