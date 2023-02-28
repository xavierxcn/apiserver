# 打包依赖阶段使用golang作为基础镜像
FROM harbor.ks.x/dockerhub/library/golang:1.19 as builder

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /apiserver/api

COPY . .

# CGO_ENABLED禁用cgo 然后指定OS等，并go build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o api
RUN chmod +x api

FROM scratch

WORKDIR /apiserver/api

# 将上一个阶段publish文件夹下的所有文件复制进来
COPY --from=builder /apiserver/api .

ENV GIN_MODE=release

EXPOSE 80

ENTRYPOINT ["./api", "serve"]
