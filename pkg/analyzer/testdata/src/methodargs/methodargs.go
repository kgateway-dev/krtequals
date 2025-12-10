package methodargs

// helper is a package-level function that compares two values.
// This simulates patterns like proto.Equal(a, b).
func helper(a, b any) bool {
	return a == b
}

// MethodArgs tests that fields passed as arguments to helper functions are detected.
type MethodArgs struct {
	Proto1  any
	Proto2  any
	Direct  string
	Missing any // want "field \"Missing\" in type \"MethodArgs\" is not used in Equals"
}

func (m *MethodArgs) Equals(other *MethodArgs) bool {
	if m == nil && other == nil {
		return true
	}
	if m == nil || other == nil {
		return false
	}
	return m.Direct == other.Direct &&
		helper(m.Proto1, other.Proto1) &&
		helper(m.Proto2, other.Proto2)
}

// NestedMethodArgs tests nested helper calls.
type NestedMethodArgs struct {
	A       any
	B       any
	C       any
	Missing any // want "field \"Missing\" in type \"NestedMethodArgs\" is not used in Equals"
}

func (n *NestedMethodArgs) Equals(other *NestedMethodArgs) bool {
	if n == nil && other == nil {
		return true
	}
	if n == nil || other == nil {
		return false
	}
	// All three fields used as method arguments
	return helper(n.A, other.A) &&
		helper(n.B, other.B) &&
		helper(n.C, other.C)
}
