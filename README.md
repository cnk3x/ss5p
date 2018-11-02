socks5代理
====

## 编译

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -ldflags "-s -w -X main.version=`cat VERSION`"
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -mod vendor -ldflags "-s -w -X main.version=`cat VERSION`"
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -mod vendor -ldflags "-s -w -X main.version=`cat VERSION`"
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -mod vendor -ldflags "-s -w -X main.version=`cat VERSION`"
```

## 直接运行
```
usage: ss5p [<flags>]

socks5代理服务器

Flags:
      --ip=IP      如果本机没有网卡没有外网ip，需要设置外网ip
  -l, --port=8080  监听的端口
  -u, --usr=""     账号
  -p, --pwd=""     密码
```

```
./ss5p -u=<usr> -p=<pwd> -i=<public_ip> -l=<port>
```

## 在容器里运行

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -ldflags "-s -w -X main.version=`cat VERSION`" -o  docker/ss5p
upx -9 docker/ss5p
docker build -t wenaiyao/ss5p .
docker run -p 18080:8080 -e usr=<usr> -e pwd=<pwd> -e ip=<public_ip> wenaiyao/ss5p
```
