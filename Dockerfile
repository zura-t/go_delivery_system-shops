FROM golang:1.20
WORKDIR /app
COPY . .
RUN go build -o main main.go

EXPOSE 8082
CMD [ "/app/main" ]