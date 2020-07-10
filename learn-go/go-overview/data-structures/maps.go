// maps.go
// Maps are similar to hashes or dictionaries in other languages.

package main

import "fmt"

func main() {
  // Create an empty map with make: map[key-type]value-type
  m := make(map[string]int)

  // Set
  m["a"] = 1
  m["b"] = 2

  // Get
  v1 := m["a"]
  fmt.Println("m: ", m)
  fmt.Println("v1", v1)

  // Removing a key-value pair
  delete(m, "a")
  fmt.Println("delete:", m)
}
