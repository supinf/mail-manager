FROM golang:1.15.0-alpine3.12 AS build-env
ENV APPDIR=$GOPATH/src/github.com/supinf/supinf-mail/app-api \
    CGO_ENABLED=0 \
    GOOS=linux
RUN apk --no-cache add gcc musl-dev git
WORKDIR $APPDIR
COPY go.mod $APPDIR/
COPY go.sum $APPDIR/
RUN go mod download
RUN go mod verify
COPY . $APPDIR/
RUN go build -ldflags "-s -w" -o /app

FROM alpine:3.12
ENV GOPATH=/go
RUN apk add --no-cache ca-certificates
COPY i18n/locale_ja-JP.ini /go/src/github.com/supinf/supinf-mail/app-api/i18n/
COPY --from=build-env /app /app
ENTRYPOINT ["/app", "--host", "0.0.0.0", "--port", "80", "--scheme", "http"]
