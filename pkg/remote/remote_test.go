package remote

import "testing"

// TestRemoteAllSignals_IncludesAllModules verifies the fixture imports all three modules from hylla-fixture-go-2.
func TestRemoteAllSignals_IncludesAllModules(t *testing.T) {
	want := "ping|HYLLA_CODEC_SENTINEL_TOKEN_V1|HYLLA_X_CLOCK_SENTINEL_TOKEN_V1"
	if got := RemoteAllSignals(); got != want {
		t.Fatalf("RemoteAllSignals() = %q, want %q", got, want)
	}
}
