FROM golang:1.20-alpine AS builder
ARG TARGETOS
ARG TARGETARCH
WORKDIR /
COPY services.account services.account
COPY keys keys 
WORKDIR /services.account
ENV CGO_ENABLED=0
COPY ./services.account/go.mod ./services.account/go.sum ./
RUN  --mount=type=cache,target=/go/pkg/mod \
    go mod download
COPY . . 
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o main ./src/cmd/main.go

FROM scratch

ENV PORT 8080

COPY --from=builder /services.account/main .
COPY --from=builder /keys ./keys
COPY --from=builder /services.account/src/locales ./src/locales

EXPOSE $PORT

CMD ["/main"]