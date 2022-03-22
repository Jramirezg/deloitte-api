# syntax=docker/dockerfile:1


FROM golang:latest AS build_base

LABEL maintainer="jramirezg.1944@gmail.com"

RUN mkdir app
ADD . /app

WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /dapi 



FROM golang:1.17-alpine

WORKDIR /app
RUN mkdir  /opt/dapi
COPY --from=build_base  /dapi .
RUN chmod +x ./dapi

EXPOSE 8080
CMD [ "./dapi" ] 