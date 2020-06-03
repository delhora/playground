FROM golang:alpine AS build-env
RUN apk --no-cache add git
ADD . /src
WORKDIR /src
RUN go build -o goapp

FROM alpine
WORKDIR /app
COPY --from=build-env /src/goapp /app/
ENTRYPOINT ./goapp