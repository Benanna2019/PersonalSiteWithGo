FROM docker.io/golang:1.23.3-alpine AS build

# Install build dependencies
RUN apk add --no-cache upx nodejs npm
RUN npm install -g pnpm

# Install Task
RUN wget -O- https://taskfile.dev/install.sh | sh

WORKDIR /src
COPY . ./

# Install dependencies and generate static files
RUN go mod download
RUN pnpm install
RUN task build:styles    # Generate CSS
RUN task build:templ    # Generate templ files
RUN task build:esbuild  # Bundle JS

# Build the Go binary
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -ldflags="-s" -o /bin/main .
RUN upx -9 -k /bin/main

FROM scratch
COPY --from=build /bin/main /
COPY --from=build /src/web/static /web/static
COPY --from=build /src/web/custom-elements /web/custom-elements
ENTRYPOINT ["/main"]