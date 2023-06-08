FROM golang

WORKDIR /app

COPY . .

RUN go get -u github.com/Coolenov/Fusion-library
RUN go get -u github.com/go-sql-driver/mysql
RUN go get github.com/joho/godotenv

CMD ["go","run", "main.go"]