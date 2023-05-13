FROM golang:1.20
LABEL authors="twoCookie"

WORKDIR /app

COPY . /app

RUN go build -o main .

EXPOSE 9999

ENTRYPOINT ["/app/main"]
