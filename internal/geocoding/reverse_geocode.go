package geocoding

import (
	"context"
	"errors"
	"github.com/kevinmichaelchen/api-geo/internal/idl/coop/drivers/geo/v1beta1"
	"googlemaps.github.io/maps"
)

var errNoResults = errors.New("no results found for coordinates")

type Place struct {
	ID      string
	Address string
	Types   []string
}

func ReverseGeocode(ctx context.Context, c *maps.Client, lat, lng float64) (*v1beta1.ReverseGeocodeResponse, error) {
	results, err := c.ReverseGeocode(ctx, &maps.GeocodingRequest{
		Address:    "",
		Components: nil,
		Bounds:     nil,
		Region:     "",
		LatLng: &maps.LatLng{
			Lat: lat,
			Lng: lng,
		},
		ResultType:   nil,
		LocationType: nil,
		PlaceID:      "",
		Language:     "",
		Custom:       nil,
	})
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errNoResults
	}
	out := new(v1beta1.ReverseGeocodeResponse)
	for _, res := range results {
		out.Results = append(out.Results, resultToProto(res))
	}
	return out, nil
}

func resultToProto(in maps.GeocodingResult) *v1beta1.GeocodingResult {
	return &v1beta1.GeocodingResult{
		AddressComponents: addressComponentsToProto(in.AddressComponents),
		FormattedAddress:  in.FormattedAddress,
		AddressGeometry:   geometryToProto(in.Geometry),
		Types:             in.Types,
		PlaceId:           in.PlaceID,
	}
}

func addressComponentsToProto(in []maps.AddressComponent) []*v1beta1.AddressComponent {
	var out []*v1beta1.AddressComponent
	for _, e := range in {
		out = append(out, addressComponentToProto(e))
	}
	return out
}

func addressComponentToProto(in maps.AddressComponent) *v1beta1.AddressComponent {
	return &v1beta1.AddressComponent{
		LongName:  in.LongName,
		ShortName: in.ShortName,
		Types:     in.Types,
	}
}

func geometryToProto(in maps.AddressGeometry) *v1beta1.AddressGeometry {
	return &v1beta1.AddressGeometry{
		Location:     latLngToProto(in.Location),
		LocationType: locationTypeToProto(in.LocationType),
		Bounds:       latLngBoundsToProto(in.Bounds),
		Viewport:     latLngBoundsToProto(in.Viewport),
		Types:        in.Types,
	}
}

func locationTypeToProto(in string) v1beta1.LocationType {
	switch in {
	case "ROOFTOP":
		return v1beta1.LocationType_LOCATION_TYPE_ROOFTOP
	case "RANGE_INTERPOLATED":
		return v1beta1.LocationType_LOCATION_TYPE_RANGE_INTERPOLATED
	case "GEOMETRIC_CENTER":
		return v1beta1.LocationType_LOCATION_TYPE_GEOMETRIC_CENTER
	case "APPROXIMATE":
		return v1beta1.LocationType_LOCATION_TYPE_APPROXIMATE
	default:
		return v1beta1.LocationType_LOCATION_TYPE_INVALID
	}
}

func latLngBoundsToProto(in maps.LatLngBounds) *v1beta1.LatLngBounds {
	return &v1beta1.LatLngBounds{
		Northeast: latLngToProto(in.NorthEast),
		Southwest: latLngToProto(in.SouthWest),
	}
}

func latLngToProto(in maps.LatLng) *v1beta1.LatLng {
	return &v1beta1.LatLng{
		Latitude:  in.Lat,
		Longitude: in.Lng,
	}
}
