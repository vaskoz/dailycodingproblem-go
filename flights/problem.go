package flights

// Airport represents the airport code.
type Airport string

// Flight from a departure airport to an arrival airport.
type Flight struct {
	depart, arrive Airport
}

// Itinerary takes a list of flights and returns the itinerary.
// Returns nil if there is no possible itinerary that uses all flights.
func Itinerary(segments []Flight, route []Airport) []Airport {
	if len(segments) == 0 {
		return route
	}
	last := route[len(route)-1]
	for i, segment := range segments {
		segmentsWoCurrent := append([]Flight{}, segments[:i]...)
		segmentsWoCurrent = append(segmentsWoCurrent, segments[i+1:]...)
		route = append(route, segment.arrive)
		if segment.depart == last {
			return Itinerary(segmentsWoCurrent, route)
		}
		route = route[:len(route)-1]
	}
	return nil
}
