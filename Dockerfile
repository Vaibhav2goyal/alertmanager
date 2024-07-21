FROM golang:1.22.3 as builder
WORKDIR /app
COPY . .
RUN GOARCH=amd64 go build -o alertmanager

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app/alertmanager /app/alertmanager
COPY ./scripts/start.sh ./scripts/start.sh
COPY ./.env ./.env
EXPOSE 8080
ENTRYPOINT ["./scripts/start.sh"]
