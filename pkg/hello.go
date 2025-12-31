// Package pkg provides minimal Go fixtures for Hylla tests.
package pkg

import "fmt"

// Greeter is a tiny fixture type.
type Greeter struct {
	// Name is the greeting target.
	Name string
}

// Greet formats a greeting for the given name.
func Greet(name string) string {
	return fmt.Sprintf("hello %s", name)
}

// Say returns a greeting based on the receiver name.
func (g Greeter) Say() string {
	if g.Name == "" {
		return Greet("world")
	}
	return Greet(g.Name)
}
