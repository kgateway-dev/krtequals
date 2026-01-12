package nestedmissing

// NestedType is a struct without an Equals method
type NestedType struct {
	Value1 int
	Value2 string
}

// ParentStruct has a field of NestedType
// If ParentStruct.Equals uses == comparison on NestedType,
// it will use Go's default struct comparison
type ParentStruct struct {
	Name   string
	Nested NestedType
}

// This Equals uses default Go struct comparison for Nested field
// The linter should ideally warn that NestedType doesn't have its own Equals
func (p ParentStruct) Equals(other ParentStruct) bool {
	return p.Name == other.Name && p.Nested == other.Nested
}

// ParentWithDelegation properly delegates to nested Equals
type ParentWithDelegation struct {
	Name   string
	Nested NestedWithEquals
}

type NestedWithEquals struct {
	Value1 int
	Value2 string
}

func (n NestedWithEquals) Equals(other NestedWithEquals) bool {
	return n.Value1 == other.Value1 && n.Value2 == other.Value2
}

func (p ParentWithDelegation) Equals(other ParentWithDelegation) bool {
	return p.Name == other.Name && p.Nested.Equals(other.Nested)
}
