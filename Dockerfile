FROM docker.io/golang:1.23.3-alpine AS build

RUN apk add --no-cache upx

WORKDIR /src
COPY . ./
RUN go mod download
RUN --mount=type=cache,target=/root/.cache/go-build \
go build -ldflags="-s" -o /bin/main .
RUN upx -9 -k /bin/main

FROM scratch
ENV PORT=9001
COPY --from=build /bin/main /
COPY --from=build /src/web/static /web/static
COPY --from=build /src/web/custom-elements /web/custom-elements
COPY --from=build /src/wasm/enhance-ssr.wasm /wasm/enhance-ssr.wasm
COPY --from=build /src/content /content
ENTRYPOINT ["/main"]