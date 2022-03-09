# 使用说明与程序结构说明


## 生产环境使用说明:

1.生成服务器端(一般是Linux环境)可执行文件;

```shell script
make sync     // 生成同步程序
make check    // 生成检查程序
```

2.服务器部署

通过scp命令将两个编译完成的文件及.env文件上传到服务器:


```shell script
scp eth-sync config.env unbuntu@<服务器外网IP>:/home/ubuntu/ethsync/
scp eth-check config.env unbuntu@<服务器外网IP>:/home/ubuntu/ethcheck/
```


3.在服务器上修改.env 文件

每个程序都可以使用属于它自己的.env配置文件, 按需修改; 具体配置说明可以看.env内的文件说明;

有以下几点注意:

* SYNC_LIMIT_CONCURRENCY 同步并发数, 最好只开CPU核数的一半, 如果有单独的运行追块程序的节点, 可以多开, 最好不要超过CPU核心数, 开太多的话, geth服务节点扛不住;
* SYNC_END_BLOCK_NUMBER 区块号终止数, 入库记录不包含当前设置的数, 比如设置成100, 实际入库的最大区块号是99, 与python range函数的原理相同;
* SYNC_SORT_STYLE 程序的执行方式, asc: 按区块号从小到大升序执行同步或检查程序; desc: 按区块号从大到小降序执行同步或检查程序;
* DB_DRY_RUN  用于只输出SQL, 不执行任何数据库操作


4.进入各自目录执行:

```shell script
cd /home/ubuntu/ethsync/
nohup ./eth-sync &


cd /home/ubuntu/ethcheck/
nohup ./eth-check &
```

执行后会在目录中生成 logs.txt 文件, 可以使用 tail logs.txt 查看处理情况.



## 目录结构说明：


```text
/
├── Dockerfile              // 同步程序的Dockfile
├── Dockerfile-check        // 检查程序的Dockfile
├── Makefile                // 快捷命令文件
├── README.md
├── Structure.md
├── check.go                // 检查程序的主文件
├── config                  // 配置文件目录
│   └── settings.go         // 主配置初始化加载文件
├── eth                     // 程序主目录
│   ├── check_data.go       // 补漏程序的主函数文件
│   ├── constants.go        // 常量定义
│   ├── entry.go            // 主程序入口函数
│   ├── number_util.go      // 处理起始块号的工具函数
│   └── sync_handler.go     // 同步程序的主函数文件
├── go.mod
├── go.sum
├── init.go                 // 程序初始执行函数
├── model                   // 模型类
│   ├── balance.go
│   ├── block.go
│   ├── contract.go
│   ├── log.go
│   ├── node.go
│   ├── token.go
│   ├── token_transfer.go
│   └── transcation.go
├── pkg                     // 工具包
│   ├── ethclient           // 从geth官方抽出来的eth客户端, 做了些小改动, 官网的客户端有些数据取不到
│   └── utils               // 工具函数
├── repo                    // Repository类包
│   ├── balance.go
│   ├── block.go
│   ├── contract.go
│   ├── log.go
│   ├── node.go
│   ├── redis.go
│   ├── token.go
│   ├── token_transfer.go
│   └── transcation.go
├── service                 // 服务中间件
│   ├── balance.go
│   ├── block.go
│   ├── contract.go
│   ├── log.go
│   ├── node.go
│   ├── token.go
│   ├── token_transfer.go
│   └── transcation.go
├── sql                     // 数据库SQL
│   └── init.sql
├── sync.go                 // 同步数据主文件
└── test
    └── eth_test.go
```


## 执行流程:

### 检查程序路径:

sync.go --> eth/entry.go --> eth/sync_handler.go

### 检查程序路径:

check.go --> eth/entry.go --> eth/check_data.go