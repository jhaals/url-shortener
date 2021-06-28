FROM golang:1.16.4 as base

ENV GO111MODULE=on
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

FROM base AS builder

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags='-s -w -extldflags "-static"' .

FROM scratch
COPY --from=builder /app/url-shortener .
COPY public public
CMD ["./url-shortener"] 
EXPOSE 1337
