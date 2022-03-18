# Kratos 接入 DTM 例子

## 版本依赖

**dtm LTS版本**
**kratos LTS版本**

## dtm server

### 代码下载

```shell
git clone git@github.com/dtm-labs/dtf.git
```

### 配置

```yaml
MicroService:
 Driver: 'dtm-driver-kratos' # name of the driver to handle register/discover
 Target: 'discovery://127.0.0.1:2379/dtmservice' # register dtm server to this url
 EndPoint: 'grpc://localhost:36790'
```

### 运行dtm

```shell
DISCOVERY_ENDPOINT=127.0.0.1:2379 go run main.go -c your_conf.yml # DISCOVERY_ENDPOINT 换成你的 etcd 服务地址
```

## 启动service

```shell
cd kratos/trans

make build && ./bin/trans -conf configs/config.yaml
```

## api层调用

```shell
cd kratos/app

DISCOVERY_ENDPOINT=127.0.0.1:2379 go run main.go # DISCOVERY_ENDPOINT 换成你的 etcd 服务地址
```