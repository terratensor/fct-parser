FROM golang:latest AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ./fct-parser .


# 2

FROM scratch

LABEL org.opencontainers.image.source = https://github.com/audetv/fct-parser

WORKDIR /app

COPY --from=build /app/fct-parser /app/fct-parser
COPY --from=build /app/parsed-files /app/parsed-files
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Europe/Moscow

EXPOSE 8080

# Create a volume for data storage
VOLUME /app/parsed-files

CMD ["./fct-parser", "-j", "-h"]
