FROM golang:1.11.1-alpine3.8 as build-env

RUN apk add --update --no-cache ca-certificates git


RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app

FROM scratch
COPY --from=build-env /go/bin/app /go/bin/app


ENTRYPOINT ["/go/bin/app"]