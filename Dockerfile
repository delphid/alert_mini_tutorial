# syntax=docker/dockerfile:1

FROM golang:1.16-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN mkdir log
RUN go build -o /my-app
EXPOSE 80
CMD [ "/my-app" ]
