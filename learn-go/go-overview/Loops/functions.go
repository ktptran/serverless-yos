// functions.go

package main

import "fmt"

// This function takes two ints and returns their sum as an int
func add(a int, b int) int {
  return a + b
}

func main() {
  result := add(1, 2)
  fmt.Println("1+2 = ", result)
}
