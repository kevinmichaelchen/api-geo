package main

import (
	"context"
	"github.com/kevinmichaelchen/api-geo/internal/distance"
	"github.com/kevinmichaelchen/api-geo/internal/geocoding"
	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
	"log"
	"os"
)

const (
	officeLat     = 40.7427627
	officeLng     = -73.936122
	restaurantLat = 40.7183895
	restaurantLng = -73.9526267
)

func main() {
	c := getClient()

	ctx := context.Background()

	geoRes1, err := geocoding.ReverseGeocode(ctx, c, officeLat, officeLng)
	if err != nil {
		log.Fatal(err)
	}
	pretty.Println(geoRes1[0])

	geoRes2, err := geocoding.ReverseGeocode(ctx, c, restaurantLat, restaurantLng)
	if err != nil {
		log.Fatal(err)
	}
	pretty.Println(geoRes2[0])

	res, err := distance.BetweenPlaces(ctx, c, geoRes1[0].ID, geoRes2[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	pretty.Println(res)
}

func getClient() *maps.Client {
	apiKey := os.Getenv("API_KEY")
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	return c
}
