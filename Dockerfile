FROM hub.hamdocker.ir/library/golang:1.20.3

WORKDIR /go/src/github.com/harleywinston/x-bot

COPY ./ .

RUN go build -buildvcs=false -o ./build/x-bot ./cmd

CMD ["./build/x-bot"]
