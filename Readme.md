# API for save and get image


## Run the app

go run main.go

### Request

`POST /saveImage` with image in FormData key="image" value= your image

### Response

if ok
HTTP/1.1 200 OK
"Image save! ID is %Id%"

if not ok
HTTP/1.1 400 OK
"Can not save image!"

### Request

`GET /saveImage?id=%id%` with id for you save image

### Response

if ok
HTTP/1.1 200 OK
Image

if not ok
HTTP/1.1 400 OK
"We have not image for this id!"