# golang-web-server-starter

This project provide a golang web server starter template and have a simple api service example using Uncle Bob's clean architecture.

## Contents

- [Overview](#Overview)
- [Run](#Run)
- [Packages Used](#Packages-Used)
- [Tools Used](#Tools-Used)

## Overview

- Go project layout
- Using [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) in web server service
- [Gin](https://github.com/gin-gonic/gin) as HTTP web framework
- [GORM](https://github.com/jinzhu/gorm) to manipulate MySQL
- [zap](https://github.com/uber-go/zap) to log data
- Database migration by using [goose](https://github.com/pressly/goose)
- Generate API document([OpenAPI Specification](https://swagger.io/specification)) from code by using [goas](https://github.com/mikunalpha/goas)

## Run

```shell
# start a mysql server run in docker container
make run-db

# database migration
make run-migrate-db

# if you encounter error like:
# [mysql] 2019/07/15 03:37:52 packets.go:36: unexpected EOF
# 2019/07/15 03:37:52 goose run: driver: bad connection
# make: *** [run-migrate-db] Error 1
# please wait a few seconds to wait mysql server running and retry database migration command again

# run web server
make
```

## Packages Used

- [Gin](https://github.com/gin-gonic/gin): HTTP web framework
- [GORM](https://github.com/jinzhu/gorm): ORM library
- [zap](https://github.com/uber-go/zap): logging library
- [yaml.v3](https://gopkg.in/yaml.v3): encode and decode YAML values

## Tools Used

- [goose](https://github.com/pressly/goose): database migration tool.
- [goas](https://github.com/mikunalpha/goas): generate [OpenAPI Specification](https://swagger.io/specification) json file with comments in Go.

## License

[MIT License](LICENSE)
