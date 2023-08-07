FROM golang:1.20 as builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o coffeeshop_app .


FROM scratch
COPY --from=builder /app ./

CMD ["./coffeeshop_app"]
