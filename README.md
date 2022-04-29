# api-geo

Proof of concept geo-server using the Google Maps API.

This could be useful for Fulfillment, where we sort/rank the drivers who are 
best able to fulfill a trip, accounting for their proximity to the trip's pickup
location.

You'll need an `API_KEY` env var which you can generate 
[here](https://console.cloud.google.com/google/maps-apis/credentials).
Also make sure you have [enabled](https://console.cloud.google.com/google/maps-apis/api-list) 
the APIs linked below.

Capabilities:
1. [Reverse geocoding](https://developers.google.com/maps/documentation/geocoding/overview): converting geographic coordinates into a Place.
2. [Distance/duration calculation](https://developers.google.com/maps/documentation/distance-matrix/distance-matrix)

# Generate protos
We generate protos with [buf](https://docs.buf.build/installation)
```
buf generate idl
```

There is a [maintained repo](https://github.com/googleapis/googleapis) of 
protobufs for Google APIs, but it doesn't contain anything for Google Maps.