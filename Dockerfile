FROM golang:1.20.1-alpine
EXPOSE 8082
WORKDIR /app
COPY . /app/
RUN go build -o SE2023
ENTRYPOINT ["/app/SE2023"]