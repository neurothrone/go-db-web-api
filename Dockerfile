FROM golang:1.23.2-alpine

COPY . .
RUN go get -d -v
RUN go build -o /app/cmd/main

# EXPOSE 8080

ENTRYPOINT [ "/app/cmd/main" ]
