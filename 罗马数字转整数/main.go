package main

import (
	"flag"
	"fmt"
)

var p string

func init(){
    flag.StringVar(&p, "v", "", "...")
    flag.Parse()
}

func main() {
    fmt.Printf("value : %v\n", p)
    fmt.Printf("result : %v\n", romanToInt(p))
}


func romanToInt(s string) int {
    romanMap := map[string]int{
        "I":1,
        "V":5,
        "X":10,
        "L":50,
        "C":100,
        "D":500,
        "M":1000,
    }

    data := []byte(s)


    var p, result int
    p = 0
    for i := len(data) -1; i >= 0; i-- {

        x := romanMap[string(data[i])]
        if x >= p {
            result +=x
        }else{
            result -=x
        }
        p = x
    }
    return result
}
