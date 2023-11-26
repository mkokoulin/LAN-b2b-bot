FROM golang:buster

WORKDIR /app
ADD . .
RUN go build -o /usr/local/bin/hello-world

RUN env

EXPOSE 8080
CMD ["/usr/local/bin/hello-world"]


