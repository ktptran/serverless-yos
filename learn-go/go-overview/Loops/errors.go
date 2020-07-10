// errors.go
package main

import "errors"
import "fmt"

// By convention, errors are the last return value and have type error,
// a built-in interface.
func f1(arg int) (int, error) {
  if args == 42 {
    return -1, errors.New("42 does not compute")
  }
  return arg + 1, nil
}
