# Build Stage
# First pull Golang image
FROM golang:1.18.1-alpine as build-env
 
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY config ./config
COPY controllers ./controllers
COPY database ./database
COPY entities ./entities
COPY main.go config.json ./

RUN CGO_ENABLED=0 go build -v -o golangcrud ./main.go

FROM alpine:3.16.2
WORKDIR /bin
COPY --from=build-env /app/golangcrud /bin/golangcrud
COPY --from=build-env /app/config.json /bin/config.json
# Add docker-compose-wait tool -------------------
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait
EXPOSE 8080
CMD /bin/golangcrud
