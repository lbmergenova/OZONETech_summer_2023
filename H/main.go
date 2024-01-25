package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type T struct {
	x     int
	count int
}

func f(m map[int][]int, k, n int) []string {
	var arr []T
	ipStr := []string{"0", "0"}
	for v := range m {
		arr = append(arr, T{v, len(m[v])})

	}
	if len(arr) > k {
		ip := "100.200.0.0/16"
		ipStr = append(ipStr, ip)
		ipStr[0] = strconv.Itoa(65536 - n)
		ipStr[1] = "1"
		return ipStr
	}
	nblist := 0
	count := k
	sort.Slice(arr, func(i, j int) bool { return arr[i].count < arr[j].count })
	for i := range arr {
		if arr[i].count+len(arr)-i-1 <= k {
			for j := range m[arr[i].x] {
				ip := "100.200." + strconv.Itoa(arr[i].x) + "." + strconv.Itoa(m[arr[i].x][j])
				ipStr = append(ipStr, ip)
				k--
			}
		} else {
			ip := "100.200." + strconv.Itoa(arr[i].x) + "." + "0/24"
			ipStr = append(ipStr, ip)
			nblist += 256 - arr[i].count
			k--
		}
	}
	count -= k
	ipStr[0] = strconv.Itoa(nblist)
	ipStr[1] = strconv.Itoa(count)
	return ipStr
}

func main() {
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	var n, k int
	var str string
	m := make(map[int][]int)
	fmt.Fscan(input, &n)
	fmt.Fscan(input, &k)
	count := n
	if n <= k {
		output.WriteString("0\n")
		output.WriteString(strconv.Itoa(n))
		output.WriteString("\n")
		for n > 0 {
			fmt.Fscan(input, &str)
			output.WriteString(str)
			output.WriteString("\n")
			n--
		}
	} else {
		for n > 0 {
			fmt.Fscan(input, &str)
			arr := strings.Split(str, ".")
			x, _ := strconv.Atoi(arr[2])
			y, _ := strconv.Atoi(arr[3])
			m[x] = append(m[x], y)
			n--
		}
		str := f(m, k, count)
		for i := range str {
			output.WriteString(str[i])
			output.WriteString("\n")
		}
	}
}
