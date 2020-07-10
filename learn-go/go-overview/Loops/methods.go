// methods.go

package main

import "fmt"

type Rect struct {
  width, height int
}

// Methods can be defined for either pointer or value receiver types.
// Here's an example of a value receiver.
func (r Rect) perim() int {
  return 2*r.width + 2*r.height
}

// This area method has a receiver type of *rect
// You may want to use a pointer receiver type to avoid copying on method
// calls or to allow the method to mutate the receiving struct
func (r *Rect) area() int {
  return r.width * r.height
}

func main() {
  // Here we call the 2 methods defined for our struct.
  // Go automatically handles the conversion between values and pointers
  // for method calls.
  r := &Rect{width: 10, height: 5}
  fmt.Println("area: ", r.area())
  fmt.Println("perim:", r.perim())
}
