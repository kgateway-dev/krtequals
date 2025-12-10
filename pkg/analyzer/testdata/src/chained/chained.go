package chained

// ChainedEquals tests that fields used in chained && expressions are detected.
type ChainedEquals struct {
	Field1  int
	Field2  string
	Field3  bool
	Missing int // want "field \"Missing\" in type \"ChainedEquals\" is not used in Equals"
}

func (c ChainedEquals) Equals(other ChainedEquals) bool {
	return c.Field1 == other.Field1 &&
		c.Field2 == other.Field2 &&
		c.Field3 == other.Field3
}

// DeeplyChained tests longer chains with mixed operators.
type DeeplyChained struct {
	A       int
	B       int
	C       int
	D       int
	E       int
	Missing int // want "field \"Missing\" in type \"DeeplyChained\" is not used in Equals"
}

func (d DeeplyChained) Equals(other DeeplyChained) bool {
	return d.A == other.A &&
		d.B == other.B &&
		d.C == other.C &&
		d.D == other.D &&
		d.E == other.E
}
