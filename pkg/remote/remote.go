// Package remote exercises cross-repo, multi-module dependency edges for fixture graphs.
package remote

import (
	"strings"

	label "github.com/evanmschultz/hylla-fixture-go-2/codec/pkg/label"
	"github.com/evanmschultz/hylla-fixture-go-2/pkg/ping"
	frozen "github.com/evanmschultz/hylla-fixture-go-2/x/clock/pkg/frozen"
)

// RemotePing returns the root-module ping response from hylla-fixture-go-2.
func RemotePing() string {
	return ping.Ping()
}

// RemoteCodecLabel returns the nested codec-module sentinel from hylla-fixture-go-2.
func RemoteCodecLabel() string {
	return label.Sentinel()
}

// RemoteClockLabel returns the nested x/clock-module sentinel from hylla-fixture-go-2.
func RemoteClockLabel() string {
	return frozen.Sentinel()
}

// RemoteAllSignals returns all cross-repo module responses in a stable order.
func RemoteAllSignals() string {
	return strings.Join([]string{
		RemotePing(),
		RemoteCodecLabel(),
		RemoteClockLabel(),
	}, "|")
}
