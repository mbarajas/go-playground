package main

import "fmt"

func main() {
  sample := "Test123"
  charArray := []rune(sample)
  charArray2 := []byte(sample)

  fmt.Printf("hello, world\n")
  fmt.Println(charArray)
  fmt.Println(charArray2)
  fmt.Println(string(charArray))
  fmt.Println(string(charArray2))
}
