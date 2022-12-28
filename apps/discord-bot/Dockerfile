FROM golang:1.19.4 as base
FROM base as prepare

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY . .
RUN go mod download && \
    go mod verify && \
    go test ./... && \
    mkdir -p /app/tmp && \
    go build -v -o /app/tmp/tardy-toad .

FROM base as run

COPY --from=prepare /app/tmp/tardy-toad /usr/local/bin/tardy-toad

CMD ["tardy-toad"]