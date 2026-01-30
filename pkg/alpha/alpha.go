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

// Hello returns a greeting for the provided name.
func Hello(name string) string {
	if strings.TrimSpace(name) == "" {
		name = "world"
	}
	return DefaultPrefix + " " + name
}

// Shout returns an uppercased greeting for the provided name.
func (g Greeter) Shout(name string) string {
	prefix := strings.TrimSpace(g.Prefix)
	if prefix == "" {
		prefix = DefaultPrefix
	}
	if strings.TrimSpace(name) == "" {
		name = "world"
	}
	return strings.ToUpper(prefix + " " + name)
}
