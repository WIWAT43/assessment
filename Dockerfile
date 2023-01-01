FROM golang:1.19

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o ./out/assessment .

CMD ["/app/out/assessment"]