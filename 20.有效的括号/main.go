package main

import (
	"flag"
	"fmt"
)

var p string

func init() {
	flag.StringVar(&p, "v", "", "...")
	flag.Parse()
}

func main() {
	fmt.Printf("value : %v\n", p)
	fmt.Printf("result : %v\n", isValid("()"))
}

func isValid(s string) bool {
	if len(s) == 0 || len(s) % 2 != 0 {
		return false
	}

	pair := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

    stack := []byte{}
	for i := 0; i < len(s); i++ {

        if pair[s[i]] > 0{
            stack = append(stack, pair[s[i]])
        }else{
            if(len(stack) == 0 || stack[len(stack)-1] != s[i]){
                return false
            }
            stack = stack[:len(stack) -1]
        }
    }

    return len(stack) == 0
}

