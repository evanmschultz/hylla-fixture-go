package beta

import "fmt"

// BranchScopeGreeting returns a deterministic string used by branch-scope ingest tests.
func BranchScopeGreeting(name string) string {
	return fmt.Sprintf("branch-scope:%s", name)
}
