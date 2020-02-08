package day440

// Polymorph is a thing that can 'DoIt'
// but specifics are based on implementation.
type Polymorph interface {
	DoIt() string
}

const (
	slothMsg       = "I'm not really feeling it."
	industriousMsg = "It's already done."
)

// Sloth can 'DoIt' in a relaxed way.
type Sloth struct {
}

// Industrious can 'DoIt' in a productive way.
type Industrious struct {
}

// DoIt from a Sloth.
func (s Sloth) DoIt() string {
	return slothMsg
}

// DoIt from Industrious.
func (i Industrious) DoIt() string {
	return industriousMsg
}

// AskPolymorph queries the poly how it does it.
func AskPolymorph(poly Polymorph) string {
	return poly.DoIt()
}
