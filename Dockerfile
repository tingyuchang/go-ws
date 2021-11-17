FROM golang:1.16-alpine

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go .
RUN go mod tidy
RUN go build -o /go-ws-app

ENV PORT=8080
EXPOSE ${NAME}
CMD ["sh", "-c", "/go-ws-app -name $NAME"]