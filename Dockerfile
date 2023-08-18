FROM cgr.dev/chainguard/go:dev-1.20
WORKDIR /app
COPY . /app
RUN go build -o myapp
ENTRYPOINT ["./myapp"]
