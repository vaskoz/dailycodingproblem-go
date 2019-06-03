package day285

// SunsetViews returns the number of building with some amount
// of sunset views.
// Buildings are ordered from east to west.
// Runs in O(N) time and O(1) space.
func SunsetViews(buildingHeights []int) int {
	if len(buildingHeights) == 0 {
		return 0
	}
	sunsetViews := 1
	shadow := buildingHeights[len(buildingHeights)-1]
	for i := len(buildingHeights) - 2; i >= 0; i-- {
		if buildingHeights[i] > shadow {
			sunsetViews++
			shadow = buildingHeights[i]
		}
	}
	return sunsetViews
}
