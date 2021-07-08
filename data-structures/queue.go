package main

import (
  "fmt"
  "container/list"
)

func main(){
  var queue []string

  queue = append(queue, "Hello ") // Enqueue
  queue = append(queue, "world!")

  for len(queue) > 0 {
    fmt.Print(queue[0]) // First element
    queue = queue[1:]   // Dequeue
  }

  queue2 := list.New()

  queue2.PushBack("Hello ") // Enqueue
  queue2.PushBack("world!")

  for queue2.Len() > 0 {
    e := queue2.Front() // First element
    fmt.Print(e.Value)

    queue2.Remove(e) // Dequeue
  }
}
