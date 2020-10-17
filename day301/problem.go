package day301

import "github.com/algds/bloom"

// BloomFilter encapsulates github.com/algds/bloom
// to conform to method signatures.
type BloomFilter struct {
	blf bloom.Filter
}

// Add a value to the set of values.
func (bf *BloomFilter) Add(value interface{}) {
	bf.blf.Add(value)
}

// Check whether a value is in the set.
func (bf *BloomFilter) Check(value interface{}) bool {
	return bf.blf.Contains(value)
}

// NewBloomFilter creates a new instance of BloomFilter.
func NewBloomFilter(m uint, hashes []bloom.Hash) *BloomFilter {
	return &BloomFilter{blf: bloom.New(m, hashes...)}
}
