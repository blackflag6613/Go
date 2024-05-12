FROM golang:bullseye
WORKDIR /opt/app
COPY . .
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go