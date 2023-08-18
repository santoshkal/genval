FROM golang:latest
WORKDIR /app
COPY . /app
RUN go build -o myapp
ENTRYPOINT ["./myapp"]
