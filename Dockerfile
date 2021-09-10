FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN make install && make generate-proto
COPY *.go ./

RUN go build -o /server

EXPOSE 8080

CMD [ "/server", "start" ]
