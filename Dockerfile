FROM cgr.dev/chainguard/*
WORKDIR /app
COPY . /app
RUN go mod download
RUN go build -o app
