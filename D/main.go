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

	var t, k, n, m int
	var str string
	fmt.Fscan(input, &t)
	for t > 0 {
		fmt.Fscan(input, &k)
		fmt.Fscan(input, &n)
		fmt.Fscan(input, &m)
		arr := make([][]byte, n)
		for i := range arr {
			arr[i] = make([]byte, m)
			fmt.Fscan(input, &str)
			for j := range str {
				arr[i][j] = byte(str[j])
			}
		}
		k--
		for k > 0 {
			for i := range arr {
				fmt.Fscan(input, &str)
				for j := range str {
					if arr[i][j] == 46 {
						arr[i][j] = byte(str[j])
					}
				}
			}
			k--
		}
		for i := range arr {
			output.WriteString(string(arr[i]))
			output.WriteString("\n")
		}
		output.WriteString("\n")
		t--
	}
}
