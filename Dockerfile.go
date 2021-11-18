ARG VERSION=1.17.3

FROM golang:${VERSION}-alpine3.14
WORKDIR /go/src/github.com/clastix/simple-app/
#
# go mod caching
#
COPY ["go.*", "."]
RUN go mod download
#
# app compilation
#
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app

FROM scratch
COPY --from=0 /go/src/github.com/clastix/simple-app/app .
EXPOSE 80
ENTRYPOINT ["/app"]
