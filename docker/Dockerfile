FROM golang:1.21-alpine

WORKDIR /app

RUN apk update && apk add git bash
RUN git clone https://github.com/HidemaruOwO/puyodeliver.git --depth 1 /app
RUN go build -ldflags="-s -w" -trimpath main.go
RUN mv main puyodeliver

EXPOSE 3000

CMD ["./puyodeliver"]
