FROM golang:alpine
WORKDIR /c/Users/Jean\ Franco/go/src/prueba
COPY . .
RUN go build -o /go/bin/prueba main.go