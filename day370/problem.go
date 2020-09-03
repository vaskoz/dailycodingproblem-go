package day370

import (
	"log"
	"sort"
)

// DeliveryType enumerates either a pickup or dropoff delivery.
type DeliveryType string

const (
	Pickup  = "pickup"
	Dropoff = "dropoff"
)

// DeliveryAction represents the data supplied to the algorithm.
type DeliveryAction struct {
	DeliveryID uint
	Timestamp  uint
	Type       DeliveryType
}

// CourierActiveTime returns the number of seconds the courier was active
// in delivering packages.
//
// Function panics if bad data is discovered e.g. mismatch of dropoff and pickups.
// Runtime is O(N log N) due to sort.
func CourierActiveTime(data []DeliveryAction) uint {
	if len(data) == 0 {
		return 0
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Timestamp < data[j].Timestamp
	})

	var answer, startActiveTime uint

	outForDelivery := make(map[uint]DeliveryAction)

	for i, d := range data {
		if i == 0 {
			if d.Type != Pickup {
				log.Panicf("mismatched. can't have Dropoff without preceding Pickup. %v", d)
			}

			startActiveTime = d.Timestamp
			outForDelivery[d.DeliveryID] = d
		} else {
			switch d.Type {
			case Pickup:
				if len(outForDelivery) == 0 {
					startActiveTime = d.Timestamp
				} else if _, found := outForDelivery[d.DeliveryID]; found {
					log.Panicf("duplicate pickup exists. %v", d)
				}
				outForDelivery[d.DeliveryID] = d
			case Dropoff:
				if _, found := outForDelivery[d.DeliveryID]; !found {
					log.Panicf("mismatched. can't have Dropoff without preceding Pickup. %v", d)
				} else {
					delete(outForDelivery, d.DeliveryID)
					if len(outForDelivery) == 0 {
						answer += (d.Timestamp - startActiveTime)
					}
				}
			default:
				log.Panicf("invalid DeliveryType. %v", d)
			}
		}
	}

	return answer
}
