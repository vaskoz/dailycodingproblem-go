package day286

// Building represents a 2D building.
type Building struct {
	Left, Right, Height int
}

// Skyline represents vertical edges of the skyline.
type Skyline struct {
	X, Height int
}

// CreateSkyline returns the skyline.
func CreateSkyline(buildings []Building) []Skyline {
	return createSkyline(buildings, 0, len(buildings)-1)
}

func createSkyline(buildings []Building, l, r int) []Skyline {
	if l == r {
		result := make([]Skyline, 2)
		result[0] = Skyline{buildings[l].Left, buildings[l].Height}
		result[1] = Skyline{buildings[l].Right, 0}
		return result
	}
	mid := (l + r) / 2
	sl := createSkyline(buildings, l, mid)
	sr := createSkyline(buildings, mid+1, r)
	return merge(sl, sr)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(sky1, sky2 []Skyline) []Skyline {
	result := make([]Skyline, 0, len(sky1)+len(sky2))
	var h1, h2, i, j int

	for i < len(sky1) && j < len(sky2) {
		if sky1[i].X < sky2[j].X {
			x1 := sky1[i].X
			h1 = sky1[i].Height
			maxh := max(h1, h2)
			result = append(result, Skyline{x1, maxh})
			i++
		} else {
			x2 := sky2[j].X
			h2 = sky2[j].Height
			maxh := max(h1, h2)
			result = append(result, Skyline{x2, maxh})
			j++
		}
	}
	for i < len(sky1) {
		result = append(result, sky1[i])
		i++
	}
	for j < len(sky2) {
		result = append(result, sky2[j])
		j++
	}
	return result
}
