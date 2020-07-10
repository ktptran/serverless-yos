// variables.go
// Declare variables using var, := syntax

package main

import "fmt"

func main() {

  // Use var to declare 1 or more variables
  var a string = "initial"
  fmt.Println(a)
  var b, c int = 1, 2
  fmt.Println(b, c)

  // Go can infer a variable's type
  var d = true
  fmt.Println(d)

  // Types have a default zero-value, for example `int` has a '0' zero value
  var e int;
  fmt.Println(e)

  // The := shorthand below is equivalent to `var f string = "short"`
  f := "short"
  fmt.Println(f)
}
