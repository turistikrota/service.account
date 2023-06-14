FROM golang:1.20-alpine AS builder
ARG TARGETOS
ARG TARGETARCH
ARG GITHUB_TOKEN

RUN apk add --no-cache git ca-certificates

ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOPRIVATE=github.com/turistikrota/service.shared TOKEN=$GITHUB_TOKEN

WORKDIR /app

RUN echo "machine github.com login turistikrota password $TOKEN" > /root/.netrc

COPY go.* ./
RUN  --mount=type=cache,target=/go/pkg/mod \
    go mod download
COPY . . 
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o main ./src/cmd/main.go



FROM scratch

ENV PORT 8080

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/main .
COPY --from=builder /app/keys ./keys
COPY --from=builder /app/src/locales ./src/locales

EXPOSE $PORT

CMD ["/main"]