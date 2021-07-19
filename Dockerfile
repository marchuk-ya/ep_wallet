FROM golang:1.16

WORKDIR /go/src/wallet_ep

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/wallet_ep .

EXPOSE 10000

CMD ["./out/wallet_ep"]