FROM golang:1.19

LABEL maintainer="Nataly Kiriukhina"

WORKDIR /build

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
#COPY go.mod go.sum ./

COPY go.mod ./

RUN go mod download && go mod verify

COPY . .

EXPOSE 8000

RUN go build -o app ./cmd/apiserver/main.go

CMD ["./app"]