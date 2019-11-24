package day369

import "math"

// Price is a stock price represented as float64.
type Price float64

// Timestamp is a Unix timestamp.
type Timestamp uint

// Datapoint is a timestamp and a price.
type Datapoint struct {
	Timestamp Timestamp
	Price     Price
}

// StockService is the defined API for stock data points.
type StockService interface {
	AddOrUpdate(Datapoint)
	Remove(Timestamp)

	Max() Price
	Min() Price
	Avg() Price
}

type stockService struct {
	max, min, avg Price
	n             uint
	data          map[Timestamp]Datapoint
}

func (ss *stockService) AddOrUpdate(d Datapoint) {
	total := ss.avg * Price(ss.n)

	if prev, exists := ss.data[d.Timestamp]; exists {
		ss.data[d.Timestamp] = d
		ss.avg = (total + d.Price - prev.Price) / Price(ss.n)

		if d.Price < ss.min || d.Price > ss.max {
			ss.min, ss.max = ss.recalculateMinMax()
		}
	} else {
		if d.Price > ss.max {
			ss.max = d.Price
		}
		if d.Price < ss.min {
			ss.min = d.Price
		}
		ss.n++
		ss.avg = (total + d.Price) / Price(ss.n)
		ss.data[d.Timestamp] = d
	}
}

func (ss *stockService) Remove(t Timestamp) {
	if prev, exists := ss.data[t]; exists {
		delete(ss.data, t)
		total := ss.avg * Price(ss.n)
		ss.n--

		if ss.n == 0 {
			ss.min = math.MaxFloat64
			ss.max = -math.MaxFloat64
			ss.avg = 0

			return
		}

		ss.avg = (total - prev.Price) / Price(ss.n)

		if prev.Price == ss.min || prev.Price == ss.max {
			ss.min, ss.max = ss.recalculateMinMax()
		}
	}
}

func (ss *stockService) recalculateMinMax() (Price, Price) {
	max := Price(0)
	min := Price(math.MaxFloat64)

	for _, dp := range ss.data {
		if dp.Price < min {
			min = dp.Price
		}

		if dp.Price > max {
			max = dp.Price
		}
	}

	return min, max
}

func (ss *stockService) Max() Price {
	return ss.max
}

func (ss *stockService) Min() Price {
	return ss.min
}

func (ss *stockService) Avg() Price {
	return ss.avg
}

// NewStockService returns a new StockService for a single
// stock.
func NewStockService() StockService {
	return &stockService{
		min:  math.MaxFloat64,
		max:  -math.MaxFloat64,
		data: make(map[Timestamp]Datapoint),
	}
}
