FROM golang:1.13 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd .
COPY templates .

RUN ls -al

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-gin-demo

# final stage
FROM scratch
COPY --from=builder /app/go-gin-demo /app/
EXPOSE 8080
ENTRYPOINT ["/app/go-gin-demo"]