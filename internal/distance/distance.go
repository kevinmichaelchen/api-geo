package distance

import (
	"context"
	"googlemaps.github.io/maps"
)

func BetweenPlaces(ctx context.Context, c *maps.Client, placeID1, placeID2 string) (*maps.DistanceMatrixResponse, error) {
	return c.DistanceMatrix(ctx, &maps.DistanceMatrixRequest{
		// https://developers.google.com/maps/documentation/distance-matrix/distance-matrix#origins
		Origins: []string{
			"place_id:" + placeID1,
		},
		Destinations: []string{
			"place_id:" + placeID2,
		},
		Mode:                     "",
		Language:                 "",
		Avoid:                    "",
		Units:                    "",
		DepartureTime:            "",
		ArrivalTime:              "",
		TrafficModel:             "",
		TransitMode:              nil,
		TransitRoutingPreference: "",
	})
}
