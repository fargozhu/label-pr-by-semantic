FROM golang:1.14 as builder
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go build -v -o lable-me .

FROM alpine:latest
COPY --from=builder /app/lable-me /lable-me
ENTRYPOINT ["/lable-me"]