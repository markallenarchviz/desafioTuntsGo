FROM golang:latest as builder

WORKDIR /

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o cmd/main .

FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["build/.main"]