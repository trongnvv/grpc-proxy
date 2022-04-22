FROM golang:1.17-alpine AS build

ENV TZ=Asia/Ho_Chi_Minh

WORKDIR /app

COPY . .

RUN go build -v -o /client ./client

# Runtime
FROM golang:1.17-alpine

WORKDIR /app

COPY --from=build /client /client

ENTRYPOINT ["/client"]

