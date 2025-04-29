FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o trip-planner cmd/server/main.go


FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/trip-planner .
CMD ["./trip-planner"]