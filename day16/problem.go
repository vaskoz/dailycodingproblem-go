package day16

// OrderLog represents the behavior of an order log regardless of impl.
type OrderLog interface {
	Record(orderID interface{})
	GetLast(index int) interface{}
}

type orderLogSlice struct {
	log  []interface{}
	pos  int
	size int
}

// NewOrderLogSlice returns a new OrderLog based on a slice.
func NewOrderLogSlice(size int) OrderLog {
	return &orderLogSlice{log: make([]interface{}, size), pos: -1, size: size}
}

func (ol *orderLogSlice) Record(orderID interface{}) {
	ol.pos++
	if ol.pos >= ol.size {
		ol.pos = 0
	}
	ol.log[ol.pos] = orderID
}

func (ol *orderLogSlice) GetLast(index int) interface{} {
	i := ol.pos - (index - 1)
	if i < 0 {
		i += ol.size
	}
	return ol.log[i]
}
