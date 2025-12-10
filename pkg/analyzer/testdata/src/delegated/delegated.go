package delegated

// Inner is a type that has its own Equals method.
type Inner struct {
	Value int
}

func (i *Inner) Equals(other *Inner) bool {
	if i == nil && other == nil {
		return true
	}
	if i == nil || other == nil {
		return false
	}
	return i.Value == other.Value
}

// DelegatedEquals tests that fields used via delegated .Equals() calls are detected.
type DelegatedEquals struct {
	Name    string
	Inner1  *Inner
	Inner2  *Inner
	Missing *Inner // want "field \"Missing\" in type \"DelegatedEquals\" is not used in Equals"
}

func (d *DelegatedEquals) Equals(other *DelegatedEquals) bool {
	if d == nil && other == nil {
		return true
	}
	if d == nil || other == nil {
		return false
	}
	return d.Name == other.Name &&
		d.Inner1.Equals(other.Inner1) &&
		d.Inner2.Equals(other.Inner2)
}

// MixedDelegation tests a mix of direct comparison and delegated Equals.
type MixedDelegation struct {
	ID      int
	Config  *Inner
	Data    *Inner
	Missing string // want "field \"Missing\" in type \"MixedDelegation\" is not used in Equals"
}

func (m *MixedDelegation) Equals(other *MixedDelegation) bool {
	if m == nil && other == nil {
		return true
	}
	if m == nil || other == nil {
		return false
	}
	return m.ID == other.ID &&
		m.Config.Equals(other.Config) &&
		m.Data.Equals(other.Data)
}
