package main

import (
  "fmt"
)

func main(){
  var twoD [2][3]int
  for i := 0; i < 2; i++ {
    for j := 0; j < 3; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2d: ", twoD)

  /* an array with 5 rows and 2 columns*/
  var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
  var i, j int

  /* output each array element's value */
  for  i = 0; i < 5; i++ {
    for j = 0; j < 2; j++ {
      fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
    }
  }

  b := [3][4]int{
    {0, 1, 2, 3},   /*  initializers for row indexed by 0 */
    {4, 5, 6, 7},   /*  initializers for row indexed by 1 */
    {8, 9, 10, 11},   /*  initializers for row indexed by 2 */
  }

  fmt.Println(b)
}
