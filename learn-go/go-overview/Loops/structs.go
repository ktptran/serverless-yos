// structs.go
// Primary data structure used to encapsulate business logic

package main

import "fmt"

type Person struct {
  name string
  age int
}

func main() {
  // Create a new struct
  fmt.Println(Person{name: "Alice", age: 21})

  // Omitted fields will be zero-valued
  fmt.Println(Person{name: "Bob"})

  // Getters and setters use dot notation
  s := Person{name: "Ann", age: 22}
  fmt.Println(s.name)

  s.age = 30
  fmt.Println(s.age)
}
