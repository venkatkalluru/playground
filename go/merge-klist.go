package main

import(
    "fmt"
)

var print = fmt.Println

func main() {

    sortedList := [][]int{
                    {1,2,5},
                    {3, 6,10,13, 15},
                    {9, 10, 12},
                    {4,7,11},
                    {8, 21,25,99},
                  }
    finalList := mergeKSortedLists(sortedList)
    for _, v := range finalList {
        print(v)
    }
}

func mergeTwoList(list1 []int, list2 []int) []int {

    if len(list1) == 0 {
        return list2
    }
    if len(list2) == 0 {
        return list1
    }

    var i,j=0,0
    mergedList := make([]int, 0, len(list1)+len(list2))

    for i < len(list1) && j < len(list2) {
        
        if list1[i] <= list2[j] {
            mergedList = append(mergedList, list1[i])
            i++
        } else {
            mergedList = append(mergedList, list2[j])
            j++
        }
    }

    for i < len(list1) {
        mergedList = append(mergedList, list1[i])
        i++
    }

    for j < len(list2) {
        mergedList = append(mergedList, list2[j])
        j++
    }

    return mergedList
}

func mergeKSortedLists(sortedList [][]int) []int {

    if len(sortedList) == 0 {
        return []int{}
    }
    if len(sortedList) == 1 {
        return sortedList[0]
    }

    if len(sortedList) == 2 {
        return mergeTwoList(sortedList[0], sortedList[1])
    }

    length := len(sortedList)    

    return mergeTwoList(mergeKSortedLists(sortedList[0:length/2]), 
                        mergeKSortedLists(sortedList[length/2:length]))
}
