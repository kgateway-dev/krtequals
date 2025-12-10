package unexportedon

// UnexportedFields has only unexported fields.
// When checkUnexported is enabled, missing unexported fields should be flagged.
type UnexportedFields struct {
	name    string
	value   int
	missing bool // want "field \"missing\" in type \"UnexportedFields\" is not used in Equals"
}

func (u *UnexportedFields) Equals(other *UnexportedFields) bool {
	if u == nil && other == nil {
		return true
	}
	if u == nil || other == nil {
		return false
	}
	return u.name == other.name && u.value == other.value
}

// MixedFields has both exported and unexported fields.
type MixedFields struct {
	Name            string
	internal        int
	missingInternal int    // want "field \"missingInternal\" in type \"MixedFields\" is not used in Equals"
	Missing         string // want "field \"Missing\" in type \"MixedFields\" is not used in Equals"
}

func (m *MixedFields) Equals(other *MixedFields) bool {
	if m == nil && other == nil {
		return true
	}
	if m == nil || other == nil {
		return false
	}
	return m.Name == other.Name && m.internal == other.internal
}

// AllFieldsUsed should pass with no errors when all fields are compared.
type AllFieldsUsed struct {
	name  string
	value int
	extra bool
}

func (a *AllFieldsUsed) Equals(other *AllFieldsUsed) bool {
	if a == nil && other == nil {
		return true
	}
	if a == nil || other == nil {
		return false
	}
	return a.name == other.name && a.value == other.value && a.extra == other.extra
}

