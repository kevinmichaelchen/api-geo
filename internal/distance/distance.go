package distance

import (
	"context"
	"github.com/kevinmichaelchen/api-geo/internal/idl/coop/drivers/geo/v1beta1"
	"google.golang.org/protobuf/types/known/durationpb"
	"googlemaps.github.io/maps"
)

func BetweenPlaces(ctx context.Context, c *maps.Client, placeID1, placeID2 string) (*v1beta1.DistanceMatrixResponse, error) {
	res, err := c.DistanceMatrix(ctx, &maps.DistanceMatrixRequest{
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
	if err != nil {
		return nil, err
	}
	return &v1beta1.DistanceMatrixResponse{
		OriginAddresses:      res.OriginAddresses,
		DestinationAddresses: res.DestinationAddresses,
		Rows:                 rowsToProto(res.Rows),
	}, nil
}

func rowsToProto(in []maps.DistanceMatrixElementsRow) []*v1beta1.DistanceMatrixElementsRow {
	var out []*v1beta1.DistanceMatrixElementsRow
	for _, e := range in {
		out = append(out, rowToProto(e))
	}
	return out
}

func rowToProto(in maps.DistanceMatrixElementsRow) *v1beta1.DistanceMatrixElementsRow {
	return &v1beta1.DistanceMatrixElementsRow{
		Elements: elementsToProto(in.Elements),
	}
}

func elementsToProto(in []*maps.DistanceMatrixElement) []*v1beta1.DistanceMatrixElement {
	var out []*v1beta1.DistanceMatrixElement
	for _, e := range in {
		out = append(out, elementToProto(e))
	}
	return out
}

func elementToProto(in *maps.DistanceMatrixElement) *v1beta1.DistanceMatrixElement {
	return &v1beta1.DistanceMatrixElement{
		Status:            statusToProto(in.Status),
		Duration:          durationpb.New(in.Duration),
		DurationInTraffic: durationpb.New(in.DurationInTraffic),
		Distance:          distanceToProto(in.Distance),
	}
}

func distanceToProto(in maps.Distance) *v1beta1.Distance {
	return &v1beta1.Distance{
		Text:   in.HumanReadable,
		Meters: int32(in.Meters),
	}
}

func statusToProto(in string) v1beta1.DistanceMatrixElementStatus {
	switch in {
	case "OK":
		return v1beta1.DistanceMatrixElementStatus_DISTANCE_MATRIX_ELEMENT_STATUS_OK
	case "INVALID_REQUEST":
		return v1beta1.DistanceMatrixElementStatus_DISTANCE_MATRIX_ELEMENT_STATUS_INVALID_REQUEST
	case "MAX_ELEMENTS_EXCEEDED":
		return v1beta1.DistanceMatrixElementStatus_DISTANCE_MATRIX_ELEMENT_STATUS_MAX_ELEMENTS_EXCEEDED
	case "MAX_DIMENSIONS_EXCEEDED":
		return v1beta1.DistanceMatrixElementStatus_DISTANCE_MATRIX_ELEMENT_STATUS_MAX_DIMENSIONS_EXCEEDED
	case "OVER_DAILY_LIMIT":
		return v1beta1.DistanceMatrixElementStatus_DISTANCE_MATRIX_ELEMENT_STATUS_OVER_DAILY_LIMIT
	case "OVER_QUERY_LIMIT":
		return v1beta1.DistanceMatrixElementStatus_DISTANCE_MATRIX_ELEMENT_STATUS_OVER_QUERY_LIMIT
	case "REQUEST_DENIED":
		return v1beta1.DistanceMatrixElementStatus_DISTANCE_MATRIX_ELEMENT_STATUS_REQUEST_DENIED
	case "UNKNOWN_ERROR":
		return v1beta1.DistanceMatrixElementStatus_DISTANCE_MATRIX_ELEMENT_STATUS_UNKNOWN_ERROR
	default:
		return v1beta1.DistanceMatrixElementStatus_DISTANCE_MATRIX_ELEMENT_STATUS_INVALID
	}
}
