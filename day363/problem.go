package day363

type curry struct {
	total int
	add   bool
}

func (c *curry) __(n int) *curry {
	if c.add {
		c.total += n
	} else {
		c.total -= n
	}
	c.add = !c.add
	return c
}

func (c *curry) Execute() int {
	return c.total
}

// AddSubtract alternately adds and subtracts curried arguments.
// It performs calculations as it curries.
// Calling Execute() returns the last computed value.
func AddSubtract(n int) *curry {
	return &curry{n, true}
}
