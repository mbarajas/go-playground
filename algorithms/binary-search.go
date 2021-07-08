package main

import (
  "fmt"
)

func main(){
  items := []int{6, 8, 13, 17, 24, 32, 44, 57, 64, 65, 73, 84}
  isPresent, index := binarySearch(17, items, 0, len(items) - 1)

  if isPresent {
    fmt.Printf("Item was found at index %d\n", index)
  } else {
    fmt.Printf("Item was found in the list\n")
  }
}

func binarySearch(item int, items []int, start_index int, end_index int) (bool, int) {
  for start_index <= end_index {
    median := (start_index + end_index) / 2

    if items[median] == item{
      return true, median
    } else
    if items[median] < item {
      start_index = median + 1
    }else{
      end_index = median - 1
    }
  }

  if start_index == len(items) || items[start_index] != item {
    return false, -1
  }

  return true, start_index
}
