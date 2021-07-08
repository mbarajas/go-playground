package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func main(){
  file, err := os.Open("log.txt")

  check(err)

  defer file.Close()

  count_log_endpoints(file)
}

func count_log_endpoints(file *os.File){
  counter_table := make(map[string]int)
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    log := string(scanner.Text())
    log_array := strings.Fields(log)
    _, found := counter_table[log_array[6]]
    if found {
      counter_table[log_array[6]] += 1
    }else{
      counter_table[log_array[6]] = 1
    }
  }

  fmt.Println("Log Counts:", counter_table)
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}
