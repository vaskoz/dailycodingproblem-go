package day5

// PairOperator is a function that can operate on a pair and return a value.
type PairOperator func(interface{}, interface{}) interface{}

// PairFunc is a function that takes a PairOperator function as an argument and returns a value.
type PairFunc func(PairOperator) interface{}

// Cons is given by the problem.
// Arguments a and b can be anything and returns a PairFunc.
func Cons(a, b interface{}) PairFunc {
	return func(f PairOperator) interface{} {
		return f(a, b)
	}
}

// Car invokes the PairFunc to operate on the pair and return the first value in the pair.
func Car(pair PairFunc) interface{} {
	return pair(func(a, b interface{}) interface{} {
		return a
	})
}

// Cdr invokes the PairFunc to operate on the pair and return the second value in the pair.
func Cdr(pair PairFunc) interface{} {
	return pair(func(a, b interface{}) interface{} {
		return b
	})
}
