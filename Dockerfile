FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV PORT=8080

RUN go build

CMD ["./go_clean_architecture_benkyo"]