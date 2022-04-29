# api-geo

Proof of concept geo-server using the Google Maps API.

Capabilities:
1. [Reverse geocoding](https://developers.google.com/maps/documentation/geocoding/overview): converting geographic coordinates into a Place.
2. [Distance/duration calculation](https://developers.google.com/maps/documentation/distance-matrix/distance-matrix)

# Generate protos
We generate protos with [buf](https://docs.buf.build/installation)
```
buf generate
```

There is a [maintained repo](https://github.com/googleapis/googleapis) of 
protobufs for Google APIs, but it doesn't contain anything for Google Maps.