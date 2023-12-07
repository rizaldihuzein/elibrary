# Library System

## How to Run
You can start the application by these methods:
### Method 1 : Normal Go Run
Run this on the root folder of this repo
```
go run app/library.go
```
### Method 2 : Using CMake
Run this make command if you want to test, build and run
```
make full
```

## Request Sample
### 1. Get Book List
You can do this curl if you want to do a book listing
```
curl --location 'localhost:8080/book/list'
```
### 2. Submit Book Pickup Order
You can do this curl if you want to do a book pickup order
a. If you want to have idempotency enabled add a `Request-ID` header
```
curl --location 'localhost:8080/book/order' \
--header 'Request-ID: 46e40c37-f246-474a-ba39-bc880fb51022' \
--header 'Content-Type: application/json' \
--data '{
  "pickup_date" : "2023-12-07 08:00:00",
  "return_date" : "2023-12-08 08:00:00",
  "books":[
    {
        "id": "/works/OL66554W",
        "subject": "fiction",
        "title": "Pride and Prejudice",
        "authors": [
          "Jane Austen"
        ],
        "edition": "OL47044678M"
      },
      {
        "id": "/works/OL138052W",
        "subject": "fiction",
        "title": "Alice'\''s Adventures in Wonderland",
        "authors": [
          "Lewis Carroll"
        ],
        "edition": "OL31754751M"
      }
  ]
}'
```
The Response will be 
```
{
    "data": {
        "id": "9a3682bb-8ee8-4df6-bd43-723889e75ecc"
    }
}
```
b. But If you omit the Request-ID then there will be no idempotency check, normally we should omit this but since this is a sample, we do it anyway
```
curl --location 'localhost:8080/book/order' \
--header 'Content-Type: application/json' \
--data '{
  "pickup_date" : "2023-12-07 08:00:00",
  "return_date" : "2023-12-08 08:00:00",
  "books":[
    {
        "id": "/works/OL66554W",
        "subject": "fiction",
        "title": "Pride and Prejudice",
        "authors": [
          "Jane Austen"
        ],
        "edition": "OL47044678M"
      },
      {
        "id": "/works/OL138052W",
        "subject": "fiction",
        "title": "Alice'\''s Adventures in Wonderland",
        "authors": [
          "Lewis Carroll"
        ],
        "edition": "OL31754751M"
      }
  ]
}'
```
The Response will be 
```
{
    "data": {
        "id": "9a3682bb-8ee8-4df6-bd43-723889e75ecc"
    }
}
```
### 3. Get Book order info
You can do this curl if you want to get a book pickup order info
```
curl --location 'localhost:8080/order?id=9a3682bb-8ee8-4df6-bd43-723889e75ecc'
```
Note : Use id from the previous Book Order Submit