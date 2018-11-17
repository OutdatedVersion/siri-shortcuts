# Build static executable
FROM golang:alpine AS build

RUN adduser -D -g '' app

RUN apk update && apk add --no-cache git ca-certificates make

COPY . $GOPATH/src/com/outdatedversion/siri-shortcuts/
WORKDIR $GOPATH/src/com/outdatedversion/siri-shortcuts/

RUN make installDependencies

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /tmp/server

# Build executing image
FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd

COPY --from=build /tmp/server /usr/bin/app/server

USER app

EXPOSE 8080

CMD ["/usr/bin/app/server"]
