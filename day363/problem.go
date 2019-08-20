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

func AddSubtract(n int) *curry {
	return &curry{n, true}
}
