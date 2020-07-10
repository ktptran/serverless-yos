// switch.go
// Switch statements express conditionals across many branches.

package main

import "fmt"
import "time"

func main() {
  // A basic switch
  i := 2
  fmt.Print("Write ", i, " as ")
  switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  case 3:
    fmt.Println("three")
  case 4, 5, 6: // Branches with multiple clauses
    fmt.Println("either four, five, or six")
  default: // A wildcard branch
    fmt.Println("something else")
  }

  // Expression can be non-constraints
  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("It's before noon")
  default:
    fmt.Println("It's after noon")
  }

  // A type switch lets you discover the types of a value
  whatAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
      fmt.Println("I'm a bool")
    case int:
      fmt.Println("I'm an int")
    default:
      fmt.Printf("Don't know type %T\n", t)
    }
  }
  whatAmI(true)
  whatAmI(1)
  whatAmI("hey")
}
