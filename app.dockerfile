FROM golang:1.22.0-alpine

ENV TZ=Asia/Jakarta
WORKDIR /app

COPY go.mod go.sum ./
RUN rm -rf vendor/* bin/*

RUN go clean -mod=mod
RUN go mod tidy
RUN go mod download && go mod verify
RUN go mod vendor

COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/library-backend cmd/library_app/http/main.go

EXPOSE 8889
ENTRYPOINT ["./bin/library-backend"]
CMD ["-config=/app/files/yml/library.docker.yml"]