package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func gift(arr []int, m int) map[int][]int {
	gift := make(map[int][]int)
	k := 0
	for i := range arr {
		if k <= arr[i] {
			k = arr[i] + 1
		}
		if k > m {
			return nil
		}
		gift[arr[i]] = append(gift[arr[i]], k)
		k++
	}
	return gift
}

func main() {
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	var n, m int
	fmt.Fscan(input, &n)
	fmt.Fscan(input, &m)
	arr := make([]int, n)
	sortArr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(input, &arr[i])
		sortArr[i] = arr[i]
	}
	sort.Ints(sortArr)
	g := gift(sortArr, m)
	if g == nil {
		output.WriteString("-1")
	} else {
		for i := range arr {
			output.WriteString(strconv.Itoa(g[arr[i]][0]))
			g[arr[i]] = g[arr[i]][1:]
			output.WriteString(" ")
		}
	}
	output.WriteString("\n")
}
