# ego 接入 DTM 例子

## 版本依赖

**dtm LTS版本**
**ego LTS版本**

## dtm server

### 代码下载

```shell
git clone git@github.com/dtm-labs/dtm.git
```

### 配置

```yaml
MicroService:
 Driver: 'dtm-driver-ego' # name of the driver to handle register/discover
 Target: 'etcd://127.0.0.1:2379/dtmservice' # register dtm server to this url
 EndPoint: 'localhost:36790'
```

### 运行dtm

```shell
go run main.go -c your_conf.yml
```

## 启动service

```shell
cd ego/trans

export EGO_DEBUG=true EGO_NAME=hello && go run main.go --config=config.toml
```

## api层调用

```shell
cd ego/app

export EGO_DEBUG=true EGO_NAME=hello-client && go run main.go --config=config.toml
```
当您在trans的日志中看到类似的日志
```shell
2022-09-17 16:30:34     INFO    elog/elog_api.go:22     [info]                                  {"desc": "transfer out 30 cents from 1"}
2022-09-17 16:30:34     INFO    egrpc/interceptor.go:325        access                                  {"comp": "server.egrpc", "compName": "server.grpc", "type": "unary", "code": 0, "ucode": 200, "desc": "OK", "event": "normal", "method": "/busi.Busi/TransOut", "cost": 0.528, "peerName": "", "peerIp": "127.0.0.1", "tid": "f141e6e80e7c40a6a13b14865d596a53"}
2022-09-17 16:30:34     INFO    elog/elog_api.go:22     [info]                                  {"desc": "transfer in 30 cents to 2"}
2022-09-17 16:30:34     INFO    egrpc/interceptor.go:325        access                                  {"comp": "server.egrpc", "compName": "server.grpc", "type": "unary", "code": 0, "ucode": 200, "desc": "OK", "event": "normal", "method": "/busi.Busi/TransIn", "cost": 0.28, "peerName": "", "peerIp": "127.0.0.1", "tid": "8b0bf21cfa3de824ba0ca3a7c6640fd1"}
```
那就是事务正常了