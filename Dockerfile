FROM golang:latest
LABEL maintainer="zaky zaky.hermawan9615@gmail.com"
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENV PORT 8000
RUN go build
RUN find . -name "*.go" -type f -delete
EXPOSE $PORT
CMD ["./gintoki"]
