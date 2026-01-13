# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./ # No dependencies yet
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o todoist_export .

# Runtime stage
FROM gcr.io/distroless/static-debian12

COPY --from=builder /app/todoist_export /todoist_export

ENV PORT=8080
EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/todoist_export"]
