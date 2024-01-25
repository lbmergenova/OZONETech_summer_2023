package main

import (
	"bufio"
	"encoding/json"
	"os"
	"strconv"
)

type inputJson struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Parent int    `json:"parent"`
}

type outputJson struct {
	Id   int          `json:"id"`
	Name string       `json:"name"`
	Next []outputJson `json:"next,omitempty"`
}

func putElem(elem *outputJson, m map[int][]outputJson) {
	if val, ok := m[elem.Id]; ok {
		elem.Next = append(elem.Next, val...)
		for i := range elem.Next {
			putElem(&elem.Next[i], m)
		}
	}
}

func getOutputJson(v []inputJson) outputJson {
	var val outputJson
	m := make(map[int][]outputJson, 0)
	for i := range v {
		if v[i].Id == 0 && v[i].Parent == 0 {
			val = outputJson{Name: v[i].Name, Id: v[i].Id}
		} else {
			elem := outputJson{Name: v[i].Name, Id: v[i].Id}
			m[v[i].Parent] = append(m[v[i].Parent], elem)
		}
	}
	putElem(&val, m)
	return val
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	var t, n int
	var rez []outputJson
	scanner.Scan()
	scan := scanner.Text()
	t, _ = strconv.Atoi(scan)
	for t > 0 {
		var jStr string
		scanner.Scan()
		scan := scanner.Text()
		n, _ = strconv.Atoi(scan)
		for n > 0 {
			scanner.Scan()
			line := scanner.Text()
			jStr += line
			n--
		}
		var v []inputJson
		json.Unmarshal([]byte(jStr), &v)
		rez = append(rez, getOutputJson(v))
		t--
	}
	buff, _ := json.Marshal(rez)
	output.WriteString(string(buff))
	output.WriteString("\n")
}
