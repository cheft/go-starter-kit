# go-api

## run
```shell
go run .
```

## compile
```shell
go build -ldflags '-w -s' -o app .
upx app
```

## docker
```shell
GOOS=linux go build -ldflags '-w -s' -o app .
upx app
docker build -t go-api .
docker run -it -p 8080:8080 go-api
```
