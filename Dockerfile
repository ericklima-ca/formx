FROM golang:1.18-alpine as builder
ENV CGO_ENABLED=0
WORKDIR /app
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .
RUN go build -o /formx

FROM gcr.io/distroless/base-debian11
LABEL maintainer="Erick Amorim <github.com/ericklima-ca>"
LABEL maintainer="Daniel Andrade <github.com/danliima>"
COPY --from=builder /mailmango /mailmango
EXPOSE 8080
ENTRYPOINT ["/mailmango"]
