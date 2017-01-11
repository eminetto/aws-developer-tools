FROM golang:alpine
ENV GOPATH /go
COPY src/main.go /go/src/main.go
RUN go build -o aws-developer-tools src/main.go

EXPOSE 8082

CMD ["/go/aws-developer-tools"]
