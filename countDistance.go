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

func Distance(lt1, ln1, lt2, ln2 float64) (km float64) {
	lat1 := degreesToRadians(lt1)
	lon1 := degreesToRadians(ln1)
	lat2 := degreesToRadians(lt2)
	lon2 := degreesToRadians(ln2)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	mi := c * earthRadiusMi
	km = mi * 1.609

	return km
}
