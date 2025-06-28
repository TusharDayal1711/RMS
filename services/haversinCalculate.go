package services

import "math"

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371
	dLat := toRadians(lat2 - lat1)
	dLon := toRadians(lon2 - lon1)
	lat1 = toRadians(lat1)
	lat2 = toRadians(lat2)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func toRadians(deg float64) float64 {
	return deg * math.Pi / 180
}
