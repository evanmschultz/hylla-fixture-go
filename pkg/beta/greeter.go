package beta

import (
	"fmt"
	"strings"

	"github.com/evanmschultz/hylla-fixture-go/pkg/alpha"
)

// FormatGreeting returns a formatted greeting using the alpha package.
func FormatGreeting(name string) string {
	return FormatGreetingWithPrefix(name, "beta")
}

// FormatGreetingWithPrefix returns a formatted greeting with a custom prefix.
func FormatGreetingWithPrefix(name, prefix string) string {
	cleanPrefix := strings.TrimSpace(prefix)
	if cleanPrefix == "" {
		cleanPrefix = "beta"
	}
	return fmt.Sprintf("[%s] %s", cleanPrefix, alpha.Hello(alpha.NormalizeName(name)))
}
