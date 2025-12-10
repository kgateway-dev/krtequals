package unexportedoff

// UnexportedFields has only unexported fields.
// When checkUnexported is disabled (default), no errors should be reported.
type UnexportedFields struct {
	name    string
	value   int
	missing bool
}

func (u *UnexportedFields) Equals(other *UnexportedFields) bool {
	if u == nil && other == nil {
		return true
	}
	if u == nil || other == nil {
		return false
	}
	// Only comparing name and value, missing is not used.
	// But since checkUnexported is off, no error should be reported.
	return u.name == other.name && u.value == other.value
}

// MixedFields has both exported and unexported fields.
type MixedFields struct {
	Name     string
	internal int
	Missing  string // want "field \"Missing\" in type \"MixedFields\" is not used in Equals"
}

func (m *MixedFields) Equals(other *MixedFields) bool {
	if m == nil && other == nil {
		return true
	}
	if m == nil || other == nil {
		return false
	}
	// Only exported Name is checked when checkUnexported is off
	return m.Name == other.Name && m.internal == other.internal
}
