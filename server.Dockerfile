FROM golang:1.17-alpine AS build

ENV TZ=Asia/Ho_Chi_Minh

WORKDIR /app

COPY . .

RUN go build -v -o /server ./server

# Runtime
FROM golang:1.17-alpine

WORKDIR /app

COPY --from=build /server /server

ENTRYPOINT ["/server"]

