FROM golang:1.23 AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian12
WORKDIR /root/
COPY --from=build /go/bin/app /
COPY input/messages.json ./input/messages.json
CMD ["/app"]