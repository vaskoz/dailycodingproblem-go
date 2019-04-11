package day232

import "testing"

func TestPrefixMapSum(t *testing.T) {
	t.Parallel()
	pms := NewPrefixMapSum()
	pms.Insert("columnar", 3)
	if sum := pms.Sum("col"); sum != 3 {
		t.Errorf("Sum should be 3 here got %d", sum)
	}
	pms.Insert("column", 2)
	if sum := pms.Sum("col"); sum != 5 {
		t.Errorf("Sum should be 5 here got %d", sum)
	}
	if sum := pms.Sum("columns"); sum != 0 {
		t.Errorf("Sum should be 0 here got %d", sum)
	}
	if sum := pms.Sum("hi there"); sum != 0 {
		t.Errorf("Sum should be 0 here got %d", sum)
	}
}

func BenchmarkPrefixMapSum(b *testing.B) {
	pms := NewPrefixMapSum()
	for i := 0; i < b.N; i++ {
		pms.Insert("columnar", 3)
		pms.Insert("column", 2)
		pms.Sum("col")
		pms.Sum("columns")
		pms.Sum("hi there")
	}
}
