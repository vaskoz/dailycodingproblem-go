package day370

import (
	"testing"
)

// nolint
var testcases = []struct {
	data       []DeliveryAction
	activeTime uint
}{
	{nil, 0},
	{
		[]DeliveryAction{ // overlapping deliveries
			{1, 10, DeliveryType("pickup")},
			{1, 20, DeliveryType("dropoff")},
			{2, 15, DeliveryType("pickup")},
			{3, 17, DeliveryType("pickup")},
			{3, 22, DeliveryType("dropoff")},
			{2, 30, DeliveryType("dropoff")},
		},
		20,
	},
	{
		[]DeliveryAction{ // non-overlapping deliveries
			{1, 10, DeliveryType("pickup")},
			{1, 20, DeliveryType("dropoff")},
			{2, 40, DeliveryType("pickup")},
			{3, 100, DeliveryType("pickup")},
			{3, 110, DeliveryType("dropoff")},
			{2, 50, DeliveryType("dropoff")},
		},
		30,
	},
}

func TestCourierActiveTime(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if activeTime := CourierActiveTime(tc.data); activeTime != tc.activeTime {
			t.Errorf("Expected %d, got %d", tc.activeTime, activeTime)
		}
	}
}

func TestCourierActiveTimeDropOffBeforePickup(t *testing.T) {
	t.Parallel()

	data := []DeliveryAction{
		{1, 1573280047, DeliveryType("pickup")},
		{1, 1570320725, DeliveryType("dropoff")},
		{2, 1570321092, DeliveryType("pickup")},
		{3, 1570321212, DeliveryType("pickup")},
		{3, 1570322352, DeliveryType("dropoff")},
		{2, 1570323012, DeliveryType("dropoff")},
	}

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic due to the first event by time being a Dropoff")
		}
	}()

	CourierActiveTime(data)
}

func TestCourierActiveTimeDuplicatePickup(t *testing.T) {
	t.Parallel()

	data := []DeliveryAction{
		{3, 30, DeliveryType("pickup")},
		{3, 40, DeliveryType("pickup")},
	}

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic due to a duplicate pickup for the same delivery ID")
		}
	}()

	CourierActiveTime(data)
}

func TestCourierActiveTimeDropoffBeforePickup(t *testing.T) {
	t.Parallel()

	data := []DeliveryAction{
		{1, 10, DeliveryType("pickup")},
		{1, 20, DeliveryType("dropoff")},
		{3, 30, DeliveryType("dropoff")},
	}

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic due to a dropoff only. no pickup.")
		}
	}()

	CourierActiveTime(data)
}

func TestCourierActiveTimeInvalidDeliveryType(t *testing.T) {
	t.Parallel()

	data := []DeliveryAction{
		{3, 30, DeliveryType("pickup")},
		{4, 40, DeliveryType("gonnapanic")},
	}

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic due to a duplicate pickup for the same delivery ID")
		}
	}()

	CourierActiveTime(data)
}

func BenchmarkCourierActiveTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CourierActiveTime(tc.data)
		}
	}
}
