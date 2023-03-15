FROM golang:1.19.0

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o bin/main ./cmd/link

ENTRYPOINT ["/build/bin/main"]