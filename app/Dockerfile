FROM golang

WORKDIR /go/src/app
COPY . .

RUN go build main.go

EXPOSE 8080
CMD ["./main"]
