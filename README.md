# Flash Sale

An example of gin contains many useful features


## Installation
```
$ go get github.com/mkaykisiz/flash-sale-case
```

## How to run

### Required

- Postgres
- Redis

### Ready

Create a **init database** and import [SQL](https://github.com/mkaykisiz/flash-sale-case/blob/master/docs/sql/init.sql)

### Conf

You should modify `conf/app.ini`

```
[database]
User = postgres
Password = postgres
Host = 127.0.0.1:3306
Name = flash_sale

[redis]
Host = 127.0.0.1:6379
Password =
MaxIdle = 30
MaxActive = 30
IdleTimeout = 200
...
```

### Run
```
$ cd $GOPATH/src/flash-sale-case

$ go run main.go 
$ make review
```
#### OR
```
$ cd $GOPATH/src/flash-sale-case

$ docker-compose up
```


Project information and existing API

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
[GIN-debug] POST   /auth                     --> flash_sale/routers/api.GetAuth (3 handlers)
[GIN-debug] POST   /api/v1/products          --> flash_sale/routers/api/v1.AddProduct (4 handlers)
[GIN-debug] PUT    /api/v1/products/:id      --> flash_sale/routers/api/v1.EditProduct (4 handlers)
[GIN-debug] POST   /api/v1/flash-sales       --> flash_sale/routers/api/v1.AddFlashSale (4 handlers)
[GIN-debug] GET    /api/v1/flash-sales       --> flash_sale/routers/api/v1.GetFlashSales (4 handlers)
[GIN-debug] POST   /api/v1/flash-sales/:id/buy --> flash_sale/routers/api/v1.BuyFlashSale (4 handlers)
[GIN-debug] GET    /api/v1/flash-sales/:id   --> flash_sale/routers/api/v1.GetFlashSale (4 handlers)
[GIN-debug] PUT    /api/v1/flash-sales/:id   --> flash_sale/routers/api/v1.EditFlashSale (4 handlers)
[GIN-debug] DELETE /api/v1/flash-sales/:id   --> flash_sale/routers/api/v1.DeleteFlashSale (4 handlers)
2024/09/19 13:05:54 [info] start http server listening :8000

```
Swagger doc
http://localhost:8000/swagger/index.html

Example user:
```
username: admin
password: admin

```
## Features

- Gorm
- Swagger
- logging
- Jwt-go
- Gin
- App configurable
- Cron
- Redis
- Postgres