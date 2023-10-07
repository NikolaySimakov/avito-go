# user-segmentation-service

[![GoDoc](https://godoc.org/github.com/lib/pq?status.svg)](https://pkg.go.dev/github.com/lib/pq?tab=doc)

It is required to implement a service that stores the user and the segments to which he belongs (creating, changing, deleting segments, as well as adding and deleting users in the segment).

To start the service:

- Rename .env.example to .env, specify all the necessary variables
- Run the file `migrations/create_tables.sql` to create all the necessary tables.

Launch the application:

```
go run cmd/main.go
```

## Usage

Basic HTTP requests to the server:

### POST: /segment/ - add segment

```json
{
	"slug": "AVITO_VOICE_MESSAGES"
}
```

The server will return the same response.

### DELETE: /segment/ - delete segment

```json
{
	"slug": "AVITO_VOICE_MESSAGES"
}
```

### POST: /user/

For adding and removing users to segments.

```json
{
	"user_id": "1000",
	"add_segments": ["SEGMENT_1", "SEGMENT_2"],
	"delete_segments": ["SEGMENT_3"]
}
```

### GET: /user/

Returns user segments.

```json
{
	"user_id": "1000"
}
```

Response:

```json
["SEGMENT_1", "SEGMENT_2"]
```

### POST: /user/

An example of using the time limiter from the second task. The user with this ID will be added to the specified segment with TTL:

```json
{
	"user_id": "1000",
	"add_segments": ["SEGMENT_1"],
	"delete_segments": [],
	"ttl": 2
}
```

The deadline in this JSON request is set to 2 minutes. To make sure your data disappears after the expiration date, use:

GET: /user/

```json
{
	"user_id": "1000"
}
```

## Endpoints

- POST: `/segment/` - adding segment by slug
- DELETE: `/segment/` - delete segment by slug

- GET: `/user/` - get all segments by userId
- POST: `/user/` - add and delete user from segments

## TODO

- [x] Segment creation method. Accepts a slug (name) of a segment.
- [x] Segment removal method. Accepts a slug (name) of a segment.
- [x] Method for adding a user to a segment. Receives a list of slugs (names of segments that need to be added to the user, a list of slugs (names) of segments that need to be removed from the user, user id.
- [x] Method for obtaining active user segments. Accepts user id as input.

### ✅ Additional task 2

There are situations when we need to add a user to an experiment for a limited period. For example, give a discount for only 2 days.

Task: implement the ability to set TTL (time for automatically removing a user from a segment)

Example: We want the user to be in a segment for 2 days - for this, in the method for adding segments to the user, we pass the time the user was removed from the segment as a separate field
