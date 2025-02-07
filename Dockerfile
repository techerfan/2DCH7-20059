FROM golang:alpine
WORKDIR /app
COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go build -o output main.go
CMD [ "./output" ]