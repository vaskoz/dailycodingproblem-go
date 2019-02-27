package day188

// MakeFunctionsOriginal is the problem version.
func MakeFunctionsOriginal() []func() int {
	var funcList []func() int
	for _, i := range []int{1, 2, 3} {
		funcList = append(funcList, func() int {
			return i
		})
	}
	return funcList
}

// MakeFunctionsCorrect is the corrected version.
func MakeFunctionsCorrect() []func() int {
	var funcList []func() int
	for _, i := range []int{1, 2, 3} {
		wrapped := func(i int) func() int {
			return func() int {
				return i
			}
		}
		funcList = append(funcList, wrapped(i))
	}
	return funcList
}
