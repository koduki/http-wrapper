FROM golang:1.16-buster as builder

WORKDIR /app


ADD ./ ./
RUN go mod download

COPY . ./

RUN go build -v -o main

# Use the official Debian slim image for a lean production container.
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/main /app/main

# Run the web service on container startup.
CMD ["/app/main"]
