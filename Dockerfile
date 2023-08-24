# Stage: 0
FROM cgr.dev/chainguard/* as build
WORKDIR /app
ENV env=1
SHELL ["bash -c"]
WORKDIR /app
COPY . /app
RUN go mod download \
    && go mod verify
COPY /app /app/
ENV env=1
RUN apt update

