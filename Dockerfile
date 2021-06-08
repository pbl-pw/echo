FROM golang:alpine AS build

WORKDIR /app

COPY . .

RUN go build -o echo

FROM alpine

WORKDIR /app

COPY --from=build /app/echo .

ENTRYPOINT [ "/app/echo" ]

EXPOSE 80 443

ENV PORTS=:80,:443
