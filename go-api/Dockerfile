# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
# stage 1: build up executable "app"
FROM golang:1.20.4-alpine3.18 AS builder
WORKDIR /usr/src/app
COPY go.mod go.sum main.go ./
RUN go mod download && go mod verify
COPY config/ config/
COPY docs/ docs/
COPY routes/ routes/
COPY utils/ utils/
COPY gogovan/ gogovan/
RUN go build -v -o /usr/local/bin/app

# stage 2: generate docker image for deployment
FROM alpine:3.18.0  
# RUN apk --no-cache add ca-certificates
WORKDIR /usr/src/app
COPY --from=builder /usr/local/bin/app /usr/local/bin/app
COPY config/ config/
COPY docs/ docs/
COPY execute.sh /usr/src/app/
RUN addgroup docker && adduser -S gogovan -G docker && chown gogovan:docker /usr/local/bin/app && mkdir /usr/src/app/logs && chown gogovan:docker -R /usr/src/app/* && chmod a+x /usr/src/app/execute.ash
USER gogovan
CMD ["/usr/src/app/execute.sh"]

