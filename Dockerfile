FROM golang:alpine AS build-stage

WORKDIR /rate-limiter

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /rate-limiter/bin

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /app

COPY --from=build-stage /rate-limiter/bin /app/bin
COPY --from=build-stage /rate-limiter/conf /app/conf

ENTRYPOINT [ "/app/bin" ]