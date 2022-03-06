package main

import (
	"flag"
	"fmt"
	"strconv"
)

type arrayInt []int

func (a *arrayInt) String() string {
	return "my array int"
}

func (a *arrayInt) Set(value string) error {
	i, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	*a = append(*a, i)
	return nil
}

func main() {
	var target int
	var nums arrayInt
	flag.IntVar(&target, "t", 9, "")
	flag.Var(&nums, "l", "")
	flag.Parse()
    fmt.Printf("target : %v\n", target)
    fmt.Printf("list : %v\n", nums)
    fmt.Printf("result : %v", twoSum(nums, target))
	//fmt.Printf("result : %v", twoSum([]int{3,2,4}, 6))
}

func twoSum(nums []int, target int) []int {
	var m map[int]int = make(map[int]int)
	for i, v := range nums {
		if p, ok := m[target-v]; ok {
			return []int{p, i}
		}
		m[v] = i
	}
	return nil
}
