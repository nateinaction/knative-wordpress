FROM --platform=$BUILDPLATFORM golang:1.24 AS wp-configurator
ARG TARGETOS TARGETARCH

WORKDIR /app

# Copy the Go source code
COPY . .

# Build the Go application
ENV CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH}
RUN go build -o /app/wp-configurator main.go

FROM wordpress:php8.4-fpm

# Install relay
COPY --from=mlocati/php-extension-installer /usr/bin/install-php-extensions /usr/local/bin/
RUN install-php-extensions relay

# Copy the wp-configurator binary
COPY --from=wp-configurator /app/wp-configurator /usr/local/bin/wp-configurator

ENTRYPOINT [ "wp-configurator" ]
CMD ["php-fpm"]
