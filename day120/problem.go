package day120

// singleton represents any singleton object.
type singleton struct {
	t int
}

// global contains only number of call state.
type global struct {
	evenOrOdd int
	even, odd *singleton
}

// Global represents a global context in which singletons exist.
type Global interface {
	GetInstance() interface{}
}

// New returns a new Global context.
func New() Global {
	return &global{evenOrOdd: 0, even: &singleton{0}, odd: &singleton{1}}
}

// GetInstance returns one of two possible singletons.
// The returned singleton alternates between calls.
func (g *global) GetInstance() interface{} {
	g.evenOrOdd = (g.evenOrOdd + 1) % 2
	if g.evenOrOdd == 0 {
		return g.even
	}
	return g.odd
}
