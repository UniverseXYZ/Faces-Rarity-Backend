# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /rarity

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["/rarity-backend", "run"]

