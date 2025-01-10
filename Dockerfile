FROM golang:1.23.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /gameserver ./cmd/gameserver/main.go

FROM scratch

COPY --from=builder /gameserver /gameserver

CMD ["/gameserver"]
