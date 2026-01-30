package beta

import (
	"fmt"

	"github.com/evanmschultz/hylla-fixture-go/pkg/alpha"
)

// FormatGreeting returns a formatted greeting using the alpha package.
func FormatGreeting(name string) string {
	return fmt.Sprintf("[beta] %s", alpha.Hello(name))
}
