# Stage: 0
FROM cgr.dev/chainguard/*
WORKDIR /app
ENV test=1
COPY . /app
RUN go mod download
RUN go build -o app
SHELL ["bash -c"]

# Stage: 1
FROM scratch
WORKDIR /app
COPY /app/app /app
ENTRYPOINT ["./app"]

