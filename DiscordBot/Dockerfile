FROM golang:1.27-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bot .

FROM alpine:3.22
RUN apk add --no-cache ca-certificates
COPY --from=build /bot /bot
USER 65532:65532
CMD ["/bot"] 