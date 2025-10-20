FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o grpc-auth ./cmd/sso

FROM alpine:3.20

WORKDIR /app

COPY --from=build /app/grpc-auth .
COPY /config/config.yaml ./config.yaml

ENTRYPOINT [ "/app/grpc-auth" ] 