// range.go
// Range to iterate over elements in data structures

package main

import "fmt"

func main() {
  // iterating over a slice
  nums := []int{2, 3, 4}
  sum := 0
  for index, num := range nums {
    fmt.Println("current index: ", index)
    sum += num
  }
  fmt.Println("sum: ", sum)

  // Iterating over a map's key/value pairs
  kvs := map[string]string{"a": "apple", "b": "banana"}
  for k, v := range kvs {
    fmt.Printf("%s -> %s\n", k, v)
  }
}
