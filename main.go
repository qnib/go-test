package main

import (
	"fmt"

	"github.com/qnib/go-test/lib"
)

func stupidMin(x,y int) int {
    if x<y {
        return x
    }
    return y
}

func main() {
    fmt.Printf("stupidMin(2,3) > %d", stupidMin(2,3))
    fmt.Println(basics.Truth())
}
