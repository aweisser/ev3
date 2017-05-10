# ev3

## Precondition

You have a EV3 Brick with ev3dev and GoEV3 (https://mattrajca.github.io/GoEV3/) installed.
As pre build tarball with ARMv5 support I used http://dave.cheney.net/paste/go1.4.2.linux-arm~armv5-1.tar.gz.
This works out of the box with the EV3 Brick.

## Cross Compile with Docker

Follow instructions on http://www.ev3dev.org/docs/tutorials/using-docker-to-cross-compile/

List Docker Images
```
docker images
```

Run new Docker container
```
docker run --rm -it -v $GOPATH\src\:/src -w /src ev3cc
```

Install go 1.4.2 inside Docker container
```
cd /usr/local
sudo wget https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz
sudo tar -xzf go1.4.2.linux-amd64.tar.gz
sudo rm -f go1.4.2.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

Create Build
```
go build ev3.go
```