package main

import (
  "fmt"
)

func main(){
  var myArray [3]string = [3]string{}
  intArray := [5]int{}

  myArray[0] = "test"
  intArray[0] = 23

  fmt.Println("len:", len(intArray))

  for i:=0; i <= len(intArray) - 1; i++{
    intArray[i] = i + 2
  }

  fmt.Println(intArray)

}
