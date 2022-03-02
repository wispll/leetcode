package main

import (
	"flag"
	"fmt"
	"sort"
)

type arrayString []string

func (a *arrayString) String() string {
	return "my array int"
}

func (a *arrayString) Set(value string) error {
	*a = append(*a, value)
	return nil
}

var strs arrayString

func init() {
	flag.Var(&strs, "v", "")
	flag.Parse()
}

func main() {
	fmt.Printf("list : %v\n", strs)
    fmt.Printf("result : %v", longestCommonPrefix(strs))
	//fmt.Printf("result : %v", longestCommonPrefix([]string{"dog", "racecat", "cat"}))
}

func longestCommonPrefix(strs []string) string {
	sort.Slice(strs, func(i, j int) bool {
		return strs[i] < strs[j]
	})

	result := []byte{}
	frist := []byte(strs[0])
	last := []byte(strs[len(strs)-1])
	for i := 0; i < len(last); i++ {
        if i > len(frist) -1 {
            return string(result)
        }
		if frist[i] == last[i] {
			result = append(result, frist[i])
            continue
		}
        break
	}
	return string(result)
}
