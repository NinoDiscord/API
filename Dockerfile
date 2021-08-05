FROM golang:1.16-alpine AS builder
WORKDIR /
COPY . .
RUN go get
RUN go build

FROM alpine:latest
WORKDIR /opt/Nino/api
COPY --from=builder /api /opt/Nino/api
CMD ["/opt/Nino/api/api"]
