package day342

// Reduce is a function that takes in an array,
// a combining function, and an initial value
// and builds up a result by calling the combining
// function on each element of the array, left to right.
func Reduce(lst []interface{},
	reducer func(a, b interface{}) interface{},
	initial interface{}) interface{} {
	for _, v := range lst {
		initial = reducer(initial, v)
	}
	return initial
}
