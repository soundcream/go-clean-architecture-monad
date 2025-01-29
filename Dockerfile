FROM --platform=linux/amd64 golang:1.23.5-bookworm

COPY /cmd/main /app/main
COPY /conf/config.yml /app/conf/config.yml
COPY /i18n/* /app/i18n

WORKDIR /app

EXPOSE 8080

CMD ["./main"]
