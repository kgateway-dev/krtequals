package nested

import (
	"time"
)

// Inner is a nested struct that should have its own Equals method
type Inner struct {
	Value1 int
	Value2 string
}

// NestedStruct has fields of a custom struct type.
// The linter should check that these nested fields are compared properly.
type NestedStruct struct {
	Name   string
	Config Inner // want "field \"Config\" in type \"NestedStruct\" is not used in Equals"
	// +noKrtEquals reason: internal bookkeeping
	Cache map[string]string
}

// Missing comparison - should fail because Config is not compared
func (n NestedStruct) Equals(other NestedStruct) bool {
	return n.Name == other.Name
}

// ListenerPolicyIR mimics the structure from the issue
type ListenerPolicyIR struct {
	ct            time.Time
	defaultPolicy listenerPolicy
	perPortPolicy map[uint32]listenerPolicy
	// +noKrtEquals reason: When set to true, suppress source reporting metadata from
	// ListenerPolicy specific fields that are irrelevant to the (now deprecated) HTTPListenerPolicy. Remove when HTTPListenerPolicy is removed.
	NoOrigin bool
}

type listenerPolicy struct {
	proxyProtocol                 *int
	perConnectionBufferLimitBytes *uint32
	http                          *HttpListenerPolicyIr
}

type HttpListenerPolicyIr struct {
	Value string
}

// Equals for ListenerPolicyIR that doesn't compare unexported fields
// The linter should NOT flag this with default config (CheckUnexported: false)
func (l ListenerPolicyIR) Equals(other ListenerPolicyIR) bool {
	// NoOrigin is marked with +noKrtEquals so it shouldn't be required
	return true
}
