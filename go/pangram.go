/*
Go Definition
Interface: PangramDetector
Method signature: FindMissingLetters(sentence string)

Examples:

(Note that in the examples below, the double quotes should not be
considered part of the input or output strings.)

0)  "The quick brown fox jumps over the lazy dog"

Returns: ""
(This sentence contains every letter)

1)  "The slow purple oryx meanders past the quiescent canine"
Returns: "bfgjkvz"

2)  "We hates Bagginses!"
Returns: "cdfjklmopqruvxyz"

3)  ""
Returns: "abcdefghijklmnopqrstuvwxyz"
*/

package main

import (
    "bytes"
    "fmt"
    "strings"
)

var print = fmt.Println
var printf = fmt.Printf

const allAlphabets = "abcdefghijklmnopqrstuvwxyz"

func main() {

    //Valid/Success case
    case0 := "The quick brown fox jumps over the lazy dog"
    actual := FindMissingLetters(case0)
    compare(actual, "")

    case1 := "The slow purple oryx meanders past the quiescent canine"
    actual = FindMissingLetters(case1)
    expected := "bfgjkvz"
    compare(actual, expected)

    case2 := "We hates Bagginses!"
    actual = FindMissingLetters(case2)
    expected = "cdfjklmopqruvxyz"
    compare(actual, expected)

    case3 := ""
    actual = FindMissingLetters(case3)
    compare(actual, allAlphabets)
}

func compare(actual, expected string) bool {

    printf("expected = %s, actual=%s\n", expected, actual)
    if expected == actual {
        print("succeeded")
        return true
    }
    print("failed")
    return false
}


func FindMissingLetters(sentence string) string {

    charMap := map[string]bool{}
    for _, c := range sentence {
//        printf("%c ", c)
        s := string(c)
        if c >= 65 && c <= 90 {
           s = strings.ToLower(s) 
        }
        charMap[s] = true
    }

    var buffer bytes.Buffer
    for _, c := range allAlphabets {
        if ok := charMap[string(c)]; !ok {
            buffer.WriteString(string(c))
        }
    }
    print() 
    return buffer.String()
}
