FROM golang:1.19

RUN mkdir /app

WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN go build -o be11-api-deploy

CMD ["./be11-api-deploy"]

