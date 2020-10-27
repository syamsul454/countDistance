package CountDistance

import (
	"math"
)

const (
	earthRadiusMi = 3958 // radius of the earth in miles.
	earthRaidusKm = 6371 // radius of the earth in kilometers.
)

// Coord represents a geographic coordinate.
type Coord struct {
	Lat float64
	Lon float64
}

// degreesToRadians converts from degrees to radians.
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

func Distance() (km float64) {
	lat1 := degreesToRadians(-6.506584)
	lon1 := degreesToRadians(106.646979)
	lat2 := degreesToRadians(-6.506621)
	lon2 := degreesToRadians(106.647172)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	mi := c * earthRadiusMi
	km = mi * 1.609
	//jmlh = math.Round(km)
	return km
}
