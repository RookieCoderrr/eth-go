# Eth-BSC Go SQL


## 编译Docker镜像

编译前注意修改根目录 .env 文件的配置项

### 在本地测试环境生成镜像

```shell script
make dev-sync   // 生成同步程序镜像
make dev-check  // 生成检查程序镜像
```

### 在本地测试运行docker镜像

```shell script
make run-sync
make run-check
```

### 生成Linux平台镜像

```shell script
make sync
make check
```

### windows编译
```
set GOARCH=amd64
set GOOS=linux
go build -o F:/桌面/producer  cmd/producer.go
go build -o F:/桌面/consumer cmd/consumer.go
go build -o F:/桌面/check cmd/check.go
go build -o F:/桌面/repair cmd/repair.go
go build -o F:/桌面/test_speed cmd/test_speed.go
go build -o C:/Users/YNB/Desktop/repair cmd/repair.go
go build -o C:/Users/YNB/Desktop/check cmd/check.go
```

### linux编译
```
go build -o ./cmd/producer ./cmd/producer.go
go build -o ./cmd/consumer ./cmd/consumer.go
go build -o ./cmd/check ./cmd/check.go
go build -o ./cmd/repair ./cmd/repair.go
```