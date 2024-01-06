FROM golang AS builder

WORKDIR /usr/local/src

# Dependencies
COPY go.mod ./ 
RUN go mod download

# Build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/app cmd/app/main.go

FROM alpine as runner

COPY --from=builder usr/local/src/bin/app /
COPY --from=builder usr/local/src/api /api

CMD ["/app"]
