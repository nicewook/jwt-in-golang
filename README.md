
## sign up
```
curl --location --request POST http://localhost:8080/users/signup --header 'Content-Type: application/json' --data-raw '{"username": "someuser", "password": "somepass", "address": "foobatz", "first_name": "foo", "last_name": "bar", "email": "a@gmail.com"}'
{"error":"internal server error"}{"message":"user someuser created successfully"
```

## sign in
```
$ curl --location --request POST http://localhost:8080/users/signin --header 'Content-Type: application/json' --data-raw '{"username": "someuser", "password": "somepass"}'
```

## say hello
```
$ curl http://localhost:8080/hello -H 'Content-Type: application/json' -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzM2OTEzMDksImlhdCI6MTYzMzY4NzcwOSwidXNlcm5hbWUiOiJzb21ldXNlciJ9.kzdcvoyAoizAzvDBVnUUNRR_Z6dJe3ik4XPnwE5ntX0'
```
```
$ curl http://localhost:8080/hello -H 'Content-Type: application/json' -H 'Authorization: '
```
```
$ curl http://localhost:8080/hello -H 'Content-Type: application/json' -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzM2OTIxNDIsImlhdCI6MTYzMzY4ODU0MiwidXNlcm5hbWUiOiJzb21ldXNlciJ9.MHJoXw4S_uUpNnATjdkSGRea3vELDHYGzr7qOCbRPmQ'
```