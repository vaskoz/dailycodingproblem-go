package day325

import (
	"errors"
)

// Unit represents a unit of measure.
type Unit string

// Converter knows how to convert from/to different
// units of measure.
type Converter struct {
	table map[Unit]map[Unit]float64
}

// ErrConversionImpossible returns the error given when a conversion
// isn't possible.
func ErrConversionImpossible() error {
	return errConversionImpossible
}

var errConversionImpossible = errors.New("conversion is impossible between these units")

// Convert attempts to convert a value from one unit to another.
// Uses a DFS algorithm to find a path to that conversion.
func (c *Converter) Convert(val float64, from, to Unit) (float64, error) {
	visited := make(map[Unit]struct{})
	return c.convert(val, from, to, visited)
}

func (c *Converter) convert(val float64, from, to Unit, visited map[Unit]struct{}) (float64, error) {
	if factor, found := c.table[from][to]; found {
		return val * factor, nil
	}
	for nextTo, factor := range c.table[from] {
		if _, seen := visited[nextTo]; seen {
			continue
		}
		visited[nextTo] = struct{}{}
		if result, err := c.convert(val*factor, nextTo, to, visited); err == nil {
			return result, nil
		}
		delete(visited, nextTo)
	}
	return 0.0, ErrConversionImpossible()
}

// AddConversion adds a new conversion rule to the table in both directions.
func (c *Converter) AddConversion(factor float64, from, to Unit) {
	if _, exists := c.table[from]; !exists {
		c.table[from] = make(map[Unit]float64)
	}
	if _, exists := c.table[to]; !exists {
		c.table[to] = make(map[Unit]float64)
	}
	c.table[from][to] = factor
	c.table[to][from] = 1.0 / factor
}

// NewConverter returns a pointer to a new Converter.
func NewConverter() *Converter {
	return &Converter{make(map[Unit]map[Unit]float64)}
}
