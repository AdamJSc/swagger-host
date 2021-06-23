# swagger-host

## About

Install and host a Swagger definition doc using Swagger UI.

## Running locally

### Docker

* Install [Docker](https://docs.docker.com/desktop/) and [Docker Compose](https://docs.docker.com/compose/)

* From project root, run:
```bash
docker-compose up --build
```

* Visit http://localhost:8080

### Go / Mage

* Install [Mage](https://magefile.org/)

* From project root, run:

```bash
mage swagger
go run main.go
```

* Visit http://localhost:8080
