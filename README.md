# websim

## Usage

A lightweight web simulator to test various scenarios

| Method | Path | Usage |
| --- | --- | --- |
| GET |   /health  | Simply return a 200 |
| GET |  /status/:status | Return a response the the HTTP status code submitted |
| GET |  /timer/:bucketstart/:bucketend | Return a 200 after sleeping a random number of milliseconds between start and end |
| GET |  /timersize/:bucketstart/:bucketend/:mb | Similar to timer, but also add X mb of random data |
| GET |  /timersizebytes/:bucketstart/:bucketend/:bytes | Similar to timer, but also add X bytes of random data |


## Building

### Mac/Linux

0) set GOROOT environment variable
1) Install Go and Make
2) make

### Docker

0) set GOROOT environment variable
1) Install Docker, Go and Make
2) make docker


## Running

### Mac/Linux

```
./websim
```

### Docker

```
docker pull maguec/websim:latest
docker run -i -t -p 8080:8080 maguec/websim
```

## Testing

run either the docker container or the raw application binary

```
curl http://localhost:8080/health
```

---
Copyright © 2018, Chris Mague
