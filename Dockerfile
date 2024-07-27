ARG GO_VERSION=1.22.0

FROM golang:${GO_VERSION}-alpine AS builder

RUN go env -w GOPROXY=direct
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates && update-ca-certificates

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /onepiece-api-consumer .

FROM scratch AS runner

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY .env ./
COPY --from=builder /onepiece-api-consumer /onepiece-api-consumer

EXPOSE 5050
ENTRYPOINT [ "/onepiece-api-consumer" ]