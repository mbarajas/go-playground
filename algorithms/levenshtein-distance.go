package main
import "fmt"

func levenshtein(str1, str2 []rune) int {
  s1len := len(str1)
  s2len := len(str2)
  column := make([]int, len(str1)+1)

  for y := 1; y <= s1len; y++ {
    column[y] = y
  }
  for x := 1; x <= s2len; x++ {
    column[0] = x
    lastkey := x - 1
    for y := 1; y <= s1len; y++ {
      oldkey := column[y]
      var incr int
      if str1[y-1] != str2[x-1] {
        incr = 1
      }

      column[y] = minimum(column[y]+1, column[y-1]+1, lastkey+incr)
      lastkey = oldkey
    }
  }
  return column[s1len]
}

func minimum(a, b, c int) int {
  if a < b {
    if a < c {
      return a
    }
  } else {
    if b < c {
      return b
    }
  }
  return c
}

func main(){
  var str1 = []rune("Asheville")
  var str2 = []rune("Arizona")
  fmt.Println("Distance between Asheville and Arizona:",levenshtein(str1,str2))

  str1 = []rune("Python")
  str2 = []rune("Peithen")
  fmt.Println("Distance between Python and Peithen:",levenshtein(str1,str2))

  str1 = []rune("Orange")
  str2 = []rune("Apple")
  fmt.Println("Distance between Orange and Apple:",levenshtein(str1,str2))
}
