/*
Interface: ParticleAnimation
Method signature: animate(speed int, init string) []string

You may assume the following constraints:
- "speed" will be between 1 and 10 inclusive
- "init" will contain between 1 and 50 characters inclusive
- each character in "init" will be '.' or 'L' or 'R'
*/

package main

import (
    "fmt"
)


var print = fmt.Println

type direction struct {

    right bool
    left bool
}

var output []string

func main() {

    str := "LRRL.LR.LRR.R.LRRL."
    print(str)

    strWithX, dirMap := initializeMap(str)
    expect := expected(str)

    returnVal := animate(1, strWithX, expect, dirMap)

    for _, s := range returnVal {
        print(s)
    }
}

func animate(speed int, input, expected string, directionMap map[int]direction) []string {

    if input == expected {
        output = append(output, input)
        return output
    }
    
    output = append(output, input)

    newInput, newDirMap := computeNewInputDirMap(speed, input, directionMap)

    return animate(speed, newInput, expected, newDirMap)
}

func computeNewInputDirMap(speed int, input string, directionMap map[int]direction) (string, map[int]direction) {

    newDirMap := map[int]direction{}
    newInput := make([]byte, len(input))
    for i := range input {
        dir := direction{}
        newDirMap[i] =dir
        newInput[i] = '.'
    }

    length := len(input)
    for i, _ := range input {

        direction := directionMap[i]
        leftIndex, rightIndex := computeNewIndex(speed, length, i, direction)

        if leftIndex != -1 {
            dir := newDirMap[leftIndex]
            dir.left = true
            newDirMap[leftIndex] = dir
            newInput[leftIndex] = 'X'
        }

        if rightIndex != -1 {
            dir := newDirMap[rightIndex]
            dir.right = true
            newDirMap[rightIndex] = dir
            newInput[rightIndex] = 'X'
        }
    }

    return string(newInput), newDirMap
}

func computeNewIndex(speed, length, currIndex int, currDir direction) (int, int) {

    leftIndex, rightIndex := -1, -1
    if currDir.left == true {
        index := currIndex - speed
        if index >= 0 {
            leftIndex = index
        }
    }
    if currDir.right == true {
        index := currIndex + speed
        if index < length {
            rightIndex= index
        }
    }

    return leftIndex, rightIndex
}

func initializeMap(str string) (string, map[int]direction) {

    output := make([]byte, len(str))
    dirMap := map[int]direction{}
    for i := range(str) {
        if str[i] == '.' {
            output[i] = '.'
            continue
        } else {
            output[i] = 'X'
        }

        dir := direction{}
        if str[i] == 'L' {
           dir.left = true 
        } else {
            dir.right = true
        }
        dirMap[i] = dir 
    } 
    return string(output), dirMap 
}

func expected(str string) string {

    output := make([]byte, len(str))
    for i := range(str) {
        output[i] = '.'
    } 
    
    return string(output)
}

