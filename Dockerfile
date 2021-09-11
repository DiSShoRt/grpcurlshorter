FROM golang:latest

RUN GO111MODULE=auto

RUN mkdir /zadanie
RUN chmod +x .
WORKDIR /zadanie
COPY go.mod .
COPY go.sum .





RUN go mod download
COPY ./ ./
EXPOSE 5432
EXPOSE 8080
RUN go build -o grpcurlshorter  ./cmd/server/main.go
#RUN go run ./cmd/server/main.go
CMD ["./grpcurlshorter"]

