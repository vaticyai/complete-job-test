FROM golang:1.22.1 AS build
USER root
WORKDIR /app
COPY . .
RUN go get
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o syncpocalypse main.go

FROM alpine:3.14 AS run
WORKDIR /app
COPY --from=build /app/syncpocalypse /app/syncpocalypse
ENTRYPOINT [ "/app/syncpocalypse" ]
