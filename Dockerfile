FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o test-proto-field-order .

FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/test-proto-field-order .

CMD ["./test-proto-field-order"]

