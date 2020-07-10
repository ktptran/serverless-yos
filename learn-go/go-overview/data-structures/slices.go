// slices.go
// Slices in go are dynamically sized arrays

package main

import "fmt"

func main() {

  // Declare and initialize a slice
  s := make([]string, 3)
  fmt.Println("empty s:", s)

  // Set and get
  s[0] = "a"
  s[1] = "b"
  s[2] = "c"
  fmt.Println("s:", s)
  fmt.Println("s[2]", s[2])
  fmt.Println("len(s):", len(s))

  // Append
  s = append(s, "d")
  s = append(s, "e", "f")
  fmt.Println("append:", s)
}
