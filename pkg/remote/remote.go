// Package remote exercises third-party dependency edges for fixture graphs.
package remote

import "github.com/evanmschultz/hylla-fixture-go-2/pkg/ping"

// RemotePing returns the external ping response.
func RemotePing() string {
	return ping.Ping()
}
