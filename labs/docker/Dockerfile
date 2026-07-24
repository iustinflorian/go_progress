FROM golang:alpine

WORKDIR /app
COPY main.go .

ENV LOCATION=Paris

EXPOSE 8080

CMD ["go", "run", "main.go"]