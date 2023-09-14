# Capsmhoo | Soft Arch 2023 Final Project
capstone matching platform
## Running Frontend
requires Node.js
```
$ cd frontend
$ npm install
$ npm run dev
```

## Running Server
requires go 1.20+
### running a single service
```
$ cd server
$ go mod tidy
$ go run cmd/<service-name>/main.go
```
### composing all services
```
$ cd server
$ docker-compose -p capsmhoo up -d --build
```