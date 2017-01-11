# aws-developer-tools
Project to test AWS developer tools solutions


## Compile

    docker pull golang:alpine
    docker run -v `pwd`:/go golang:alpine go build -o aws-developer-tools src/main.go

## Execute

    docker run -v `pwd`:/go -p 8082:8082 golang:alpine ./aws-developer-tools 
