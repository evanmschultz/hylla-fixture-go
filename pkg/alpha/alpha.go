package alpha

import "strings"

const (
	// DefaultPrefix is the default greeting prefix.
	DefaultPrefix = "hello"
)

// Greeter is a small fixture type that formats greetings.
type Greeter struct {
	// Prefix is the greeting prefix to use.
	Prefix string
}

// NormalizeName returns a canonical greeting target for fixtures.
func NormalizeName(name string) string {
	if strings.TrimSpace(name) == "" {
		return "world"
	}
	return name
}

// Hello returns a greeting for the provided name.
func Hello(name string) string {
	return DefaultPrefix + " " + NormalizeName(name)
}

// Shout returns an uppercased greeting for the provided name.
func (g Greeter) Shout(name string) string {
	prefix := strings.TrimSpace(g.Prefix)
	if prefix == "" {
		prefix = DefaultPrefix
	}
	return strings.ToUpper(prefix + " " + NormalizeName(name))
}
