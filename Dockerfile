FROM golang:latest

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /app/flash-sale

### Stage 2
FROM scratch
COPY --from=builder /app/flash-sale /bin/flash-sale
ENTRYPOINT ["/bin/flash-sale"]
