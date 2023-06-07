FROM golang:alpine

WORKDIR /app

COPY . .
#COPY go.mod go.sum app/
#RUN go mod download
RUN go get -u github.com/Coolenov/Fusion-library
RUN go get -u github.com/go-sql-driver/mysql
RUN go get github.com/joho/godotenv

CMD ["go","run", "main.go"]