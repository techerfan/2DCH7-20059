From golang:alpine
EXPOSE 8080
WORKDIR /app
COPY . .
RUN go build -o output main.go
CMD [ "./output" ]