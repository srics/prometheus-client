# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ADD . /src
RUN cd /src && GOOS=linux go build -o promgo

# final stage
FROM alpine
RUN mkdir -p /etc/ssl/certs
COPY certs/bundle.crt /etc/ssl/certs
WORKDIR /app
COPY --from=build-env /src/promgo /app/
ENTRYPOINT ./promgo
