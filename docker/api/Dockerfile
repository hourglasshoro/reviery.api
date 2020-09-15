FROM golang:1.14-alpine as build

WORKDIR /go/memesan

COPY src/app ./src/app
COPY go.mod .
COPY go.sum .

RUN apk add --no-cache git \
  && go get github.com/oxequa/realize

WORKDIR /go/memesan/src/app

RUN go build -o app

FROM alpine as prod

WORKDIR /app

COPY --from=build /go/memesan/app .

RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go /memesan/app

CMD ["./app"]