package main

import (
	"fmt"
	"math"

	"googlemaps.github.io/maps"
)

func GeneratePath(polyline string) (string, error) {
	coordinates, err := maps.DecodePolyline(polyline)
	if err != nil || len(coordinates) < 2 {
		return "", err
	}

	max, min := findMaxAndMin(coordinates)
	coordinates = standardizePoints(coordinates, max, min)

	res := fmt.Sprintf("M%v %v ", coordinates[0].Lng, coordinates[0].Lat)
	for i, coordinate := range coordinates {
		if i == 0 {
			continue
		} else {
			res = fmt.Sprintf("%sL %v %v ", res, coordinate.Lng, coordinate.Lat)
		}
	}

	return res, nil
}

func findMaxAndMin(latlng []maps.LatLng) (maps.LatLng, maps.LatLng) {
	max := maps.LatLng{Lat: latlng[0].Lat, Lng: latlng[0].Lng}
	min := maps.LatLng{Lat: latlng[0].Lat, Lng: latlng[0].Lng}

	for _, coords := range latlng {
		if coords.Lat < min.Lat {
			min.Lat = coords.Lat
		}
		if coords.Lat > max.Lat {
			max.Lat = coords.Lat
		}
		if coords.Lng < min.Lng {
			min.Lng = coords.Lng
		}
		if coords.Lng > max.Lng {
			max.Lng = coords.Lng
		}
	}

	return max, min
}

func standardizePoints(points []maps.LatLng, max, min maps.LatLng) []maps.LatLng {
	res := make([]maps.LatLng, len(points))

	shift := maps.LatLng{Lat: max.Lat, Lng: min.Lng}
	scalar := 100 / math.Max(max.Lat-min.Lat, max.Lng-min.Lng)

	for i, coords := range points {
		res[i].Lat = math.Round(-100 * scalar * (coords.Lat - shift.Lat)) / 100
		res[i].Lng = math.Round(100 * scalar * (coords.Lng - shift.Lng) * math.Cos(math.Pi*coords.Lat/180)) / 100
	}

	return res
}