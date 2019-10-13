FROM golang as builder

WORKDIR /api
COPY . . 

CMD ["go run main.go"]