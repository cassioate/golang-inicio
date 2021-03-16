FROM golang:latest
LABEL Autor: "Cassio"
COPY . /var
WORKDIR /var
RUN go build
ENTRYPOINT go run main.go
EXPOSE 3000