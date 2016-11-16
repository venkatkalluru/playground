package main

import (
    "fmt"
)

var print = fmt.Println

func main() {

    n := 3
    left, right := 0,0
    printValidParanthesis(n, left, right,"")
}

func printValidParanthesis(n, left, right int, str string) {

    if len(str) == n*2 {
        print(str)
        return
    }

    if left < n {
        printValidParanthesis(n, left+1, right, str + "(")
    }

    if (right < left) {
        printValidParanthesis(n, left, right+1, str + ")")
    }
}
