FROM golang:alpine as build
ENV GO111MODULE=on \
        CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=amd64 \
        GOPROXY="https://goproxy.io"\
        VERSION=1.17
WORKDIR /app
COPY . .
RUN go build -o web

FROM alpine as prod
COPY --from=build /app/web .
EXPOSE 80
CMD ["./web"]
