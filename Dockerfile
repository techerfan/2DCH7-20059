FROM golang:alpine
WORKDIR /app
COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g ./delivery/httpserver/fiber/api.go -o ./docs/swagger
RUN go build -o output main.go
CMD [ "./output" ]