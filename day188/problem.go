package day188

// MakeFunctionsOriginal is the problem version.
func MakeFunctionsOriginal() []func() int {
	var funcList []func() int
	for _, i := range []int{1, 2, 3} {
		funcList = append(funcList, func() int {
			return i //nolint: scopelint
		})
	}
	return funcList
}

// MakeFunctionsCorrect is the corrected version.
func MakeFunctionsCorrect() []func() int {
	var funcList []func() int
	for _, i := range []int{1, 2, 3} {
		funcList = append(funcList, func(i int) func() int {
			return func() int {
				return i
			}
		}(i))
	}
	return funcList
}
