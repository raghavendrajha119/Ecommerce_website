FROM golang:1.20.6 AS build
# Working directory
WORKDIR /app
COPY go.mod go.sum ./
COPY main.go .
COPY . .
RUN go build -o bin .
EXPOSE 3000
ENTRYPOINT ["/app/bin"]
