FROM golang:alpine as builder
RUN apk add --no-cache git gcc libc-dev
RUN go get github.com/labstack/echo && go get github.com/stretchr/testify/assert
ADD . .
RUN go test -v
RUN go build -o main

FROM alpine
COPY --from=builder /go/main main
EXPOSE 5001
CMD ["./main"]
