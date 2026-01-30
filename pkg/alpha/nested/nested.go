package nested

import "github.com/evanmschultz/hylla-fixture-go/pkg/alpha"

// BuildGreeting returns an alpha greeting for the provided name.
func BuildGreeting(name string) string {
	return alpha.Hello(name)
}
