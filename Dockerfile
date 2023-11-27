FROM golang:alpine AS builder

WORKDIR /

COPY . .

ENV GOOS=linux

RUN GOARCH=$(uname -m) \
    go mod download \
    && go mod verify \
    && go build -ldflags "-s -w" -o helloemoji

FROM scratch

WORKDIR /

COPY --from=builder helloemoji /helloemoji

ENV PORTS=3000

EXPOSE 3000

ENTRYPOINT ["/helloemoji"]