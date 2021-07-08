package main

import (
  "fmt"
)


func search(txt string, pat string) {
  int M = pat.length()
  int N = txt.length()

  /* A loop to slide pat one by one */
  for i := 0; i <= N - M; i++ {

    /* For current index i, check for pattern
    match */
    for j := 0; j < M; j++{
      if (txt.charAt(i + j) != pat.charAt(j)){
        break
      }
    }
    if (j == M){ // if pat[0...M-1] = txt[i, i+1, ...i+M-1]
      System.out.println("Pattern found at index " + i)
    }
  }
}

func main() {
  txt := "AABAACAADAABAAABAA"
  String := "AABA"
  search(txt, pat)
}
