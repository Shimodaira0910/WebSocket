FROM golang:1.20

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo
ENV GO111MODULE=on
ENV GOPATH=

WORKDIR /app

COPY go.mod go.sum ./
COPY main.go ./
RUN go mod download

RUN go build -o main .

EXPOSE 443
CMD ["./main"]
