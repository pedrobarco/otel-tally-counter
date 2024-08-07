FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /otel-tally-counter

EXPOSE 8080

CMD ["/otel-tally-counter"]
