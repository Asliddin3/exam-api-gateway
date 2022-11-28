FROM golang:1.18-alpine3.15
RUN mkdir api
COPY . /exam-api-gateway
WORKDIR /exam-api-gateway
RUN go mod tidy
# VOLUME [ "/data" ]
RUN go build -o main cmd/main.go
CMD ./main
EXPOSE 8070