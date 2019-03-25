# Scrum poker
Scrum Poker without features. Just a poker for Scrum. If you use Scrum -- this poker will work. Only within Scrum shall you use this poker.

# Installation
```
cd docker
docker-compose up --detach
```

# Usage (obviously not everything ready yet)
Open your browser at http://127.0.0.1/ or create a session:
```
> curl -i -d '{"title":"My session A"}' -H "Content-Type: application/json" -X POST http://127.0.0.1/grooming_sessions
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Mon, 25 Mar 2019 20:56:37 GMT
Content-Length: 69

{"id":"7b4df44c-4f40-11e9-ab33-0242c0a80003","title":"My session A"}
```
Create another one:
```
> curl -i -d '{"title":"My session B"}' -H "Content-Type: application/json" -X POST http://127.0.0.1/grooming_sessions
http://127.0.0.1/grooming_sessions
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Mon, 25 Mar 2019 20:56:56 GMT
Content-Length: 69

{"id":"863ae512-4f40-11e9-ab33-0242c0a80003","title":"My session B"}
```
Get session by ID:
```
> curl -i http://127.0.0.1/grooming_sessions/7b4df44c-4f40-11e9-ab33-0242c0a80003 # Use your unique ID, of course
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Mon, 25 Mar 2019 20:57:31 GMT
Content-Length: 69

{"id":"7b4df44c-4f40-11e9-ab33-0242c0a80003","title":"My session A"}
```
List all sessions
```
> curl -i http://127.0.0.1/grooming_sessions                                     
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Mon, 25 Mar 2019 20:58:06 GMT
Content-Length: 604

[{"id":"101289b1-4f3f-11e9-ab33-0242c0a80003","title":"Session E"},{"id":"1299b975-4f3f-11e9-ab33-0242c0a80003","title":"Session F"},{"id":"2e965e98-4f3d-11e9-aaad-0242ac1b0003","title":"Session A"},{"id":"3225f03f-4f3d-11e9-aaad-0242ac1b0003","title":"Session B"},{"id":"3555dc35-4f3d-11e9-aaad-0242ac1b0003","title":"Session C"},{"id":"38b473ea-4f3d-11e9-aaad-0242ac1b0003","title":"Session D"},{"id":"3d9ec663-4f40-11e9-ab33-0242c0a80003","title":"Session 100"},{"id":"7b4df44c-4f40-11e9-ab33-0242c0a80003","title":"My session A"},{"id":"863ae512-4f40-11e9-ab33-0242c0a80003","title":"My session B"}]
```

To bring it down:
```
docker-compose down
```
