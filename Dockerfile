FROM golang:alpine

WORKDIR /app

COPY . .

RUN go get github.com/Coolenov/Fusion-library@main
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/joho/godotenv
RUN go get github.com/go-sql-driver/mysql


CMD ["go","run", "main.go"]