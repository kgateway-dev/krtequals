package problem

import "time"

// InnerType has fields with markers
type InnerType struct {
	Used int
	// +noKrtEquals reason: should not be compared
	Ignored string
}

// OuterType has a field of InnerType
type OuterType struct {
	Name  string
	Inner InnerType
}

// This Equals uses == which ignores the +noKrtEquals marker in InnerType
// This is the problem: the linter sees that Inner is used,
// but doesn't check that InnerType.Ignored has a +noKrtEquals marker
func (o OuterType) Equals(other OuterType) bool {
	return o.Name == other.Name && o.Inner == other.Inner // want "field \"Inner\" of struct type \"InnerType\" is compared using == which ignores \\+noKrtEquals markers"
}

// ListenerPolicyIR mimics the exact structure from the GitHub issue
type ListenerPolicyIR struct {
	ct            time.Time
	defaultPolicy listenerPolicy
	perPortPolicy map[uint32]listenerPolicy
	// +noKrtEquals reason: When set to true, suppress source reporting metadata
	NoOrigin bool
}

type listenerPolicy struct {
	proxyProtocol                 *int
	perConnectionBufferLimitBytes *uint32
	// +noKrtEquals reason: should not be compared
	http *HttpListenerPolicyIr
}

type HttpListenerPolicyIr struct {
	Value string
}

// This Equals incorrectly uses == for defaultPolicy comparison
// which bypasses the +noKrtEquals marker on listenerPolicy.http
func (l ListenerPolicyIR) Equals(other ListenerPolicyIR) bool {
	return l.ct == other.ct && l.defaultPolicy == other.defaultPolicy // want "field \"defaultPolicy\" of struct type \"listenerPolicy\" is compared using == which ignores \\+noKrtEquals markers"
}

// CorrectComparison shows the correct way to compare struct fields
type CorrectComparison struct {
	Name   string
	Config InnerTypeWithEquals
}

type InnerTypeWithEquals struct {
	Value int
	// +noKrtEquals reason: should not be compared
	Internal string
}

func (i InnerTypeWithEquals) Equals(other InnerTypeWithEquals) bool {
	return i.Value == other.Value
	// Internal is intentionally not compared due to +noKrtEquals marker
}

// This correctly uses .Equals() to delegate comparison
func (c CorrectComparison) Equals(other CorrectComparison) bool {
	return c.Name == other.Name && c.Config.Equals(other.Config)
}
