FROM docker.io/library/golang:1.15-alpine as builder

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o flashsale .

CMD ["/app/flashsale"]
