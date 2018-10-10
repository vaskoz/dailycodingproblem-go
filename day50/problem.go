package day50

// Operator represents the arithmetic operations.
type Operator int

const (
	// NON is a zero value of this constant.
	NON Operator = iota
	// ADD is addition
	ADD
	// SUB is subtraction
	SUB
	// MUL is multiplication
	MUL
	// DIV is division
	DIV
)

// ArithmeticTree represents either a node operator or a leaf value.
type ArithmeticTree struct {
	val         int
	op          Operator
	left, right *ArithmeticTree
}

// Calculate takes an ArithmeticTree and returns its calculated value.
// Runs in O(N) in that it only views the tree nodes once.
func Calculate(tree *ArithmeticTree) int {
	if tree.op == NON {
		return tree.val
	}
	left := Calculate(tree.left)
	right := Calculate(tree.right)
	var result int
	switch tree.op {
	case ADD:
		result = left + right
	case SUB:
		result = left - right
	case MUL:
		result = left * right
	case DIV:
		result = left / right
	}
	return result
}
