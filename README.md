#Book Sore

### Installation

1. export environmental variables, environmental variables are stored in `.env` file.
   Run `export $(cat ./.env | xargs)`


2. Run `make migrate-up` to create database migrations.

3. Run `make Run` to build and run the project on `prot` and `ip` that are sets in `.env` file.

### Architecture

Project develop based on [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
Robert C.Martin ("Uncle Bob").

### Endpoints

##### Version
```bash 
curl --location --request GET 'http://127.0.0.1:6000/version'
```
##### Sign-up
```bash 
curl --location --request POST 'http://127.0.0.1:6000/sign-up' \
--header 'Content-Type: application/json' \
--data-raw '{
"full_name":"",
"pseudonym":"",
"username":"",
"password":""
}'
```
##### Sign-in
```bash 
curl --location --request POST 'http://127.0.0.1:6000/sign-in' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"",
    "password":""
}'
```

##### Create Book
```bash 
curl --location --request POST 'http://127.0.0.1:6000/books' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOGM2ZTdhODYtNGUwYi00MDU3LWI1YzEtOTU5MWNlOWQwZThhIiwiZXhwIjoxNjIzMzMxNDk3fQ.6aAoIWgvOECERzbsspazqEanWaHPs70mIgO4jdvN2ps' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title":"",
    "description":"",
    "is_published": true
}'
```
##### Get Books
```bash 
curl --location --request GET 'http://127.0.0.1:6000/books'
```

### Project linted
project linted by `golangci-lint` has version `1.40.1`