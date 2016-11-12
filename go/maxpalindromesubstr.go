package main

import (

    "fmt"
)

var print = fmt.Println

var lo, maxLength int

func main() {

    str := "abcdedcdbabbabbac"
    print("string length: ", len(str))

    for i:=0; i<len(str)-1; i++ {
        maxPalindromeSubstr(str, i, i)  //for odd case
        maxPalindromeSubstr(str, i, i+1)    //for even case 
    }

    print("palindrome string is ", str[lo:lo+maxLength])
}

func maxPalindromeSubstr(str string, i, j int) {

    for i >=0 && j < len(str) && str[i] == str[j] {
        i--
        j++
    }

    if maxLength < j-i-1 {
        maxLength = j-i-1
        lo = i + 1
    }
}
