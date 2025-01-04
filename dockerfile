# download go depedencies
FROM golang:1.23.0-alpine3.20 AS base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
# build local development image
FROM golang:1.23.0-alpine3.20 AS dev
COPY --from=base /go/bin /go/bin
COPY --from=base /go/pkg /go/pkg
WORKDIR /app
RUN go install github.com/air-verse/air@latest
EXPOSE 6501
# build the go binary
FROM golang:1.23.0-alpine3.20 AS builder
COPY --from=base /go/bin /go/bin
COPY --from=base /go/pkg /go/pkg
WORKDIR /app
COPY . ./
EXPOSE 6501
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/account cmd/account.go
# create production image
FROM gcr.io/distroless/base-debian11 AS release
COPY --from=builder /app/bin/account /account
EXPOSE 6501
ENTRYPOINT ["/account"]
