FROM golang:alpine

WORKDIR /app

COPY . .

RUN go get github.com/Coolenov/Fusion-library
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/joho/godotenv

EXPOSE 8080