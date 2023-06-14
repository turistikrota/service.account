FROM golang:1.20-alpine AS builder
ARG TARGETOS
ARG TARGETARCH
ARG GITHUB_TOKEN

RUN apk add --no-cache git ca-certificates

ENV CGO_ENABLED=0 GOPRIVATE=github.com/turistikrota/service.shared TOKEN=$GITHUB_TOKEN

RUN git config --global url."https://${TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"


WORKDIR /app

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
COPY --from=builder /main .
COPY --from=builder /keys ./keys
COPY --from=builder /src/locales ./src/locales

EXPOSE $PORT

CMD ["/main"]