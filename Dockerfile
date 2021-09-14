FROM golang:1.17.1-buster

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY . ./
RUN go mod download
RUN bash install.sh
RUN make generate-proto
RUN go mod tidy
RUN go build -o /server

EXPOSE 8080

CMD [ "/server", "start" ]
