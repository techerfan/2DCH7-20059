FROM golang:alpine
# EXPOSE 8085
WORKDIR /app
COPY . .
RUN go build -o output main.go
CMD [ "./output" ]