ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

# Install Node.js and npm (needed for pnpm, tailwind, etc.)
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - \
    && apt-get install -y nodejs

# Install pnpm
RUN npm install -g pnpm

# Install Task
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

WORKDIR /usr/src/app

# Copy dependency files first
COPY go.mod go.sum ./
COPY package.json pnpm-lock.yaml ./

# Install dependencies
RUN go mod download && go mod verify
RUN pnpm install

# Copy the rest of the source code
COPY . .

# Run the build task
RUN task build

FROM debian:bookworm

# Copy the compiled binary
COPY --from=builder /usr/src/app/bin/main /usr/local/bin/run-app

# Copy static files
COPY --from=builder /usr/src/app/web/static /usr/local/bin/web/static

CMD ["run-app"]
