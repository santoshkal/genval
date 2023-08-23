# Stage: 0
FROM cgr.dev/chainguard/* as build
WORKDIR /app
ENV test=1
COPY . /app
RUN go mod download \
    && go mod verify \
    && go build -o app
SHELL ["bash -c"]

# Stage: 1
FROM scratch
WORKDIR /app
COPY --chown=65532:65532 --from=build  /app/app /app
ENTRYPOINT ["./app"]

