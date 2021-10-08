
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
$ curl http://localhost:8080/hello -H 'Content-Type: application/json' -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzM2OTA4MjEsImlhdCI6MTYzMzY4NzIyMSwidXNlcm5hbWUiOiJzb21ldXNlciJ9.CqhSQ7_A4ViJ2HYa45FeLJSDvBD8I-cEvOmXmQT2R3U'
```