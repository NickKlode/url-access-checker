FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o inhousead-test ./cmd/main.go

CMD ["./inhousead-test"]