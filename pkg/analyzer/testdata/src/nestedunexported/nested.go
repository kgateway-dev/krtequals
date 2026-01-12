package nestedunexported

import (
	"time"
)

// ListenerPolicyIR mimics the structure from the issue
// When CheckUnexported: true, unexported fields should be flagged
type ListenerPolicyIR struct {
	ct            time.Time                 // want "field \"ct\" in type \"ListenerPolicyIR\" is not used in Equals"
	defaultPolicy listenerPolicy            // want "field \"defaultPolicy\" in type \"ListenerPolicyIR\" is not used in Equals"
	perPortPolicy map[uint32]listenerPolicy // want "field \"perPortPolicy\" in type \"ListenerPolicyIR\" is not used in Equals"
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
// With CheckUnexported: true, this should flag the unexported fields
func (l ListenerPolicyIR) Equals(other ListenerPolicyIR) bool {
	// NoOrigin is marked with +noKrtEquals so it shouldn't be required
	return true
}
