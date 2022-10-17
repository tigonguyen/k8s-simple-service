###----------- BUILD STAGE ----------###
FROM golang:1.16-alpine as BUILD
WORKDIR /app

COPY . .
RUN go mod download

COPY main.go ./
RUN go build -o simple-service

###----------- RUNNING STAGE -----------###
FROM alpine
WORKDIR /app

EXPOSE 8000
COPY --from=BUILD /app/simple-service /app/simple-service

CMD ["/app/simple-service"]