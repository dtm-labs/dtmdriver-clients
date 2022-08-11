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
 Target: 'etcd:///127.0.0.1:2379/dtmservice' # register dtm server to this url
 EndPoint: 'grpc://localhost:36790'
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