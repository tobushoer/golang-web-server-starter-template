# golang-web-server-starter

This project provide a golang web server starter template and have a simple api service example using Uncle Bob's clean architecture.

## Contents

- [Overview](#Overview)
- [Run](#Run)
- [Packages Used](#Packages-Used)
- [Tools Used](#Tools-Used)

## Overview

* Go project layout
* Using [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) in web server service
* [Gin](https://github.com/gin-gonic/gin) as HTTP web framework
* [GORM](https://github.com/jinzhu/gorm) to manipulate MySQL
* [zap](https://github.com/uber-go/zap) to log data
* Generate API document([OpenAPI Specification](https://swagger.io/specification)) from code by using [goas](https://github.com/mikunalpha/goas)

## Run

```shell
# start a mysql server docker container and run web server
make
```

## Packages Used

* [Gin](https://github.com/gin-gonic/gin): HTTP web framework
* [GORM](https://github.com/jinzhu/gorm): ORM library
* [zap](https://github.com/uber-go/zap): logging library
* [yaml.v3](https://gopkg.in/yaml.v3): encode and decode YAML values

## Tools Used

* [goas](https://github.com/mikunalpha/goas): generate [OpenAPI Specification](https://swagger.io/specification) json file with comments in Go.

## License

[MIT License](LICENSE)
