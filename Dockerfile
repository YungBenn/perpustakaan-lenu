# FROM golang:alpine AS builder
# RUN apk update && apk add --no-cache git
# WORKDIR /app
# COPY . .
# RUN go build -o ./bin/main ./cmd/main.go

# FROM alpine
# WORKDIR /app
# COPY --from=builder /app/bin/main /app
# CMD [ "./main" ]

FROM golang
WORKDIR /app
COPY . .
RUN go build -o ./bin/main ./cmd/main.go
CMD [ "./bin.main" ]