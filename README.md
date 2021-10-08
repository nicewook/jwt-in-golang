```
$ curl --location --request POST http://localhost:8080/users/signup --header 'Content-Type: application/json' --data-raw '{"username": "someuser", "password": "somepass", "address": "foobatz", "first_name": "foo", "last_name": "bar"}'
{"error":"internal server error"}{"message":"user someuser created successfully"}
```

```
curl --location --request POST http://localhost:8080/users/signup --header 'Content-Type: application/json' --data-raw '{"username": "someuser", "password": "somepass", "address": "foobatz", "first_name": "foo", "last_name": "bar", "email": "a@gmail.com"}'
{"error":"internal server error"}{"message":"user someuser created successfully"
```