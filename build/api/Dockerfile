# dev
FROM golang:1.18.3-alpine as dev

WORKDIR /app

ENV GO111MODULE on
ENV CGO_ENABLED 0

RUN apk add tzdata
ENV TZ=Asia/Tokyo

RUN go install github.com/cespare/reflex@latest

COPY . .

CMD reflex -r '(\.go$|go\.mod)' -s go run ./cmd/api/main.go

# builder
FROM golang:1.18.3-alpine as builder

WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main cmd/api/main.go

# prod
FROM alpine as prod

WORKDIR /app

COPY --from=builder /src/main .
COPY conf/conf.yml /app/conf/conf.yml

RUN apk add tzdata
ENV TZ=Asia/Tokyo

CMD ["./main"]
