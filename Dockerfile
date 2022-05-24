# 构建阶段
FROM golang:1.18-alpine as builder

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY core/ core/
COPY middleware/ middleware/
COPY module/ module/
COPY constants/ constants/
COPY utils/ utils/
COPY static/ static/
COPY docs/ docs/
COPY main.go .

RUN go build --tags="production" -o main .

# 构建生产环境镜像
FROM alpine:latest as runtime

WORKDIR /app

COPY --from=0 /app/main .
COPY --from=0 /app/static/ static/

EXPOSE 8000

CMD ["./main"]