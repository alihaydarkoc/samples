# Rihtim Core Sample

This is a sample project to demonstrate the usage of the Rihtim Core server.


## How to run

* Install dependencies `go get ./...`
* Run server `go run main.go`
* Navigate to `http://localhost:8080` which contains root URL error


## Sample API requests

Requests will return same dummy responses all the time. You can test them with sample curl requests.

**Create:**

```
curl -i \
    -X POST \
    -d '{ "firstName": "Mustafa", "lastName": "Kemal", "email": "mustafa@kemal.com"}' \
    http://localhost:8080/users
```

**Query:**

`curl -i -X GET http://localhost:8080/users`

**Get:**

`curl -i -X GET http://localhost:8080/users/57b36760e63bce1eeb000004`

**Update:**

```
curl -i \
    -X PUT \
    -d '{ "firstName": "Mustafa Kemal", "lastName": "Atatürk" }' \
    http://localhost:8080/users/57b36760e63bce1eeb000004
```

**Delete:**

`curl -i -X DELETE http://localhost:8080/users/57b36760e63bce1eeb000004`

**A Bonus Request:**

`curl -i -X GET http://localhost:8080/wrong`

**Another Bonus Request:**

```
curl -i \
    -X GET \
    -H "If-Modified-Since: Wed, 21 Oct 2015 07:28:00 GMT" \
    http://localhost:8080/users
```
