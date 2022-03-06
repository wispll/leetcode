package main

import (
    "flag"
    "fmt"
)

var p int

func init(){
    flag.IntVar(&p, "v", 121, "...")
    flag.Parse()
}

func main() {
    fmt.Printf("value : %v\n", p)
    fmt.Println(isPalindrome(p))
}

func isPalindrome(x int) bool {
    if x < 0{
        return false
    }

    var tmp int
    b:=x
    for b > 0{
        tmp = tmp *10 + b % 10
        b/=10
    }

    if tmp!=x{
        return false
    }

    return true
}

