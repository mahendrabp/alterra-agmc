FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app
COPY . .
# RUN go mod init simple-jwt-go
# RUN go mod tidy
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/web-app ./main.go

FROM alpine:3.13
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
# COPY --from=build /go/src/app/configs/config/local.yaml /usr/bin/configs/config
EXPOSE 5302
ENTRYPOINT /go/bin/web-app --port 5302