//go:build linux && amd64
// +build linux,amd64

package alpha

// PlatformGreeting returns a platform-tagged greeting.
func PlatformGreeting(name string) string {
	return Hello(name)
}
