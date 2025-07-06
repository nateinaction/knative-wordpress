FROM wordpress:cli-php8.4 AS wp

WORKDIR /var/www/html

# Download and extract latest WordPress
RUN curl -o /tmp/wordpress.tar.gz https://wordpress.org/latest.tar.gz && \
    tar -xzf /tmp/wordpress.tar.gz -C . --strip-components=1

# Install redis-cache plugin
RUN curl -o /tmp/redis-cache.zip https://downloads.wordpress.org/plugin/redis-cache.2.5.4.zip && \
    unzip /tmp/redis-cache.zip -d wp-content/plugins/ && \
    cp wp-content/plugins/redis-cache/includes/object-cache.php wp-content/object-cache.php

# TODO: Remove the sleep binary when images can be mounted as volumes
FROM --platform=$BUILDPLATFORM golang:1.24 AS sleeper
ARG TARGETOS TARGETARCH

WORKDIR /app

# Copy the Go source code
COPY . .

# Build the Go application
ENV CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH}
RUN go build -o /app/sleeper main.go

FROM scratch

# Copy the WordPress site files
COPY --from=wp /var/www/html /tmp/wordpress

# Copy the sleeper binary
COPY --from=sleeper /app/sleeper /sleeper

# USER 65534:65534

CMD ["/sleeper"]
