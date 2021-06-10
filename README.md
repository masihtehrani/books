#Book Sore

### Prepare for run project
- copy .env.example to .env
- filled .env data
- copy .env to your environment `export $(grep -v '^#' .env | xargs -d '\n')`
- create database `books` in mysql  
- database migrations `make migrate-up`  
- run command `make run`

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