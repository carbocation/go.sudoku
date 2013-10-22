package main

import (
	"fmt"
)

func main() {
	m := map[string]string{}
	
	for _, s := range "abcdefg" {
		m[string(s)] = string(s)
	}
	fmt.Println(m)
	
	for _, s := range "abcdefg" {
		solve(string(s), cloneValues(m))
	}
	fmt.Println(m)
}

func solve(s string, m map[string]string) {
	if s == "d" {
		delete(m, s)
	}
}

//Clone the values map
func cloneValues(values map[string]string) map[string]string {
	cpyValues := make(map[string]string, len(values))
	for k, v := range values {
		cpyValues[k] = v
	}
	
	return cpyValues
}