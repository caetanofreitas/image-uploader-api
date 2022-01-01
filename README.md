# Image Uploader API
This is project is a image uploader with image compression (Only JPG, JPEG and JFIF). Live demo [here](https://ailurus.com.br).

## Instalation
### Requirements
  - Go 1.17

  or

  - Docker
  - Docker Compose

Using Docker:

```shell
$ docker-compose up --build -d
$ docker exec -it uploader_app_1 bash
$ go get
$ go run main.go
```

Without Docker:
```shell
$ go get
$ go run main.go
```

## Environment
```.env
# USING AWS
USE_AWS= // MAKE PROJECT USE AWS
S3_NAME= // NAME FOR AWS S3 IMAGE UPLOAD
S3_URL= // AWS S3 URL TO API CONNECTION
AWS_REGION= // AWS REGION
AWS_ACCESS_KEY= // AWS ACCOUNT ACCESS KEY
AWS_SECRET= // AWS ACCOUNT SECRET
DATABASE_TYPE= // DATABASE TYPE (mysql, sqlite3, etc)
DATABASE_CONNECTION= // DATABASE STRING CONNECTION

 # WITHOUT AWS (ALL LOCAL)
DATABASE_TYPE=sqlite3
DATABASE_CONNECTION=database.db
UPLOAD_URL=http://localhost:3000/uploaded
```

## Usage
Actual endpoints:
```shell
[POST] /upload - Upload image endpoint
[GET] / -  Health Check
[GET] /uploaded/{id} - If using local uploader, get the uploaded image
[GET] /images - List all uploaded images
[GET] /images/sse - SSE endpoint to update uploaded image list
[GET] /images/{id} - Get uploaded image info
```