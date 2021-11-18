FROM golang:1.16-alpine

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go .
RUN go mod tidy
RUN go build -o /go-ws-app
COPY wait-for .
RUN chmod +x ./wait-for

ENV NAME wsapp
ENV KAFKA_ADDR kafka:9092
ENV GROUPID 0
EXPOSE 8080
CMD ["./wait-for", "kafka:9092", "--", "sh", "-c", "/go-ws-app -name ${NAME} -kfkaddr ${KAFKA_ADDR} -gid ${GROUPID}"]