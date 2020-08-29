# build go stage
FROM golang:latest AS build
WORKDIR /app

# speed up
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
ENV CGO_ENABLED=0
ENV GOOS=linux

COPY . .
RUN go build -o ./server . && chmod +x ./server

# build fed stage
FROM node:latest AS buildFed

WORKDIR /app

COPY . .
RUN cd ./frontend && yarn && yarn build

# prod stage
FROM alpine:latest AS prod
WORKDIR /app

COPY --from=build /app/server /app/server
RUN mkdir -p /app/frontend/build
COPY --from=buildFed /app/frontend/build /app/frontend/build
COPY ./config.default.json /app/config.json
COPY ./statics /app/statics

EXPOSE 8888
EXPOSE 8889
RUN chmod +x ./server

ENTRYPOINT ["/app/server"]
