#Book Sore

## Installation

1. export environmental variables, environmental variables are stored in `.env` file.
   Run `export $(cat ./.env | xargs)`


2. Run `make migrate-up` to create database migrations.

3. Run `make Run` to build and run the project on `prot` and `ip` that are sets in `.env` file.

## Architecture

Project develop based on [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
Robert C.Martin ("Uncle Bob").

## Endpoints

### Version
```bash 
curl --location --request GET 'http://127.0.0.1:6000/version'
```
response
```bash
{
    "tag": "677b9e4e74162b6ff897830992f5534a86a99801",
    "commit": "677b9e4",
    "date": "2021-06-15T09:17:19+04:30",
    "service": "books"
}
```
### Sign-up
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
response
```bash
{
    "success": true
}
```
### Sign-in
```bash 
curl --location --request POST 'http://127.0.0.1:6000/sign-in' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"",
    "password":""
}'
```
response
```bash
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNDNiZTYyOWEtMmZkNC00YWQ4LTkwNTMtMzMyZGRkN2JiZmE2IiwiZXhwIjoxNjIzNzY0MTYzfQ.x1XY4YbwEfEexKPBW5q1qLq4mvIg8e-szoP2aXT2kMU",
    "expire_at": "2021-06-15T18:06:03.081707+04:30"
}
```

### Create Book
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
response
```bash
{
    "Success": true
}
```
### Get Books
```bash 
curl --location --request GET 'http://127.0.0.1:6000/books'
```
response
```bash
{
    "data": [
        {
            "id": "1b0f86b7-fb9d-4469-9b96-b5cb5f9b02f2",
            "title": "The Man Who Laughs",
            "description": "The Man Who Laughs is a novel by Victor Hugo, originally published in April 1869 under the French title L'Homme qui rit. It takes place in England in the 1680s and 1700s, during the reigns of James II and Queen Anne, respectively, and depicts England's royalty and aristocracy of the time as cruel and power-hungry.",
            "is_published": true,
            "author_full_name": "Victor Hugo",
            "author_pseudonym": "Hugo",
            "created_at": "2021-06-15T17:07:41Z"
        },
        {
            "id": "533cc88d-1ad7-42e2-930c-fc4ee5f41274",
            "title": "Bug-Jargal",
            "description": "Bug-Jargal is a novel by the French writer Victor Hugo. First published in 1826, it is a reworked version of an earlier short story of the same name published in the Hugo brothers' magazine Le Conservateur littéraire in 1820.",
            "is_published": true,
            "author_full_name": "Victor Hugo",
            "author_pseudonym": "Hugo",
            "created_at": "2021-06-15T08:46:16Z"
        }
    ],
    "meta": {
        "current_page": 1,
        "per_page": 10,
        "last_page": 1,
        "total": 2
    }
}
```
### Update Book
```bash
curl --location --request PUT 'http://127.0.0.1:6000/books/:id' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOGM2ZTdhODYtNGUwYi00MDU3LWI1YzEtOTU5MWNlOWQwZThhIiwiZXhwIjoxNjIzNjM5MDA5fQ.1xnWPG6HCtOABp0R-p9WcS7BQGrfFaIAtGE0O0FapbQ' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title":"",
    "description":"",
    "is_published": false
}'
```
response
```bash
{
    "Success": true
}
```
### Delete Book
```bash
curl --location --request DELETE 'http://127.0.0.1:6000/books/:id' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNDNiZTYyOWEtMmZkNC00YWQ4LTkwNTMtMzMyZGRkN2JiZmE2IiwiZXhwIjoxNjIzNzY0MTYzfQ.x1XY4YbwEfEexKPBW5q1qLq4mvIg8e-szoP2aXT2kMU' \
--data-raw ''
```
response
```bash
{
    "Success": true
}
```

### Project linted
project linted by `golangci-lint` has version `1.40.1`