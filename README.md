Location History Server
=======================

## Task Description

Your task is to implement a toy in-memory location history server.

Clients should be able to speak JSON over HTTP to the server. The three endpoints it should support are:
* `POST /location/{order_id}/now`
* `GET /location/{order_id}?max=<N>`
* `DELETE /location/{order_id}`

Details about the endpoints:

`POST /location/{order_id}/now` - append a location to the history for the specified order.
Example interaction:
```
POST /location/def456/now
{
	"lat": 12.34,
	"lng": 56.78
}

200 OK
```
`GET /location/{order_id}?max=<N>` - Retrieve at most N items of history for the specified order. The most recent locations (in chronological order of insertion) should be returned first, if history is truncated by the `max` parameter.
Example interaction:
```
GET /location/abc123?max=2

200 OK
{
	"order_id": "abc123",
	"history": [
		{"lat": 12.34, "lng": 56.78},
		{"lat": 12.34, "lng": 56.79}
	]
}
```
The `max` query parameter may or may not be present. If it is not present, the entire history should be returned.

`DELETE /location/{order_id}` - delete history for the specified order. Example interaction:
```
DELETE /location/xyz987

200 OK
```

## How to build and run
- Add to environment variables `LOCATION_HISTORY_TTL_SECONDS` and `HISTORY_SERVER_LISTEN_ADDR` with values of the ttl and port number respectively.
- Run `go build` from the root of the project to generate an executable file
- Run the executable file 
- Test all cases using Postman i.e
    - Get without any orders
    - Post to add location history
    - Get added location history with and without `max` query
    - Delete location and test getting