FROM golang:1.22.2-alpine3.19

WORKDIR /src/app

RUN go install github.com/cosmtrek/air@latest

# Install 1Password CLI (secret management tool)
# https://developer.1password.com/docs/cli/get-started
RUN apk add --no-cache curl \
    && curl -sS https://downloads.1password.com/linux/static/1password-latest.tar.gz | tar -xz -C /usr/local/bin --strip-components=1

COPY . .

RUN go mod tidy


COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]