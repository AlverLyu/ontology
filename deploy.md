# 联盟链部署说明

version 0.1

## 1. 准备节点钱包

假设节点有如下私钥和证书：

```
.
├── pri
└── cert.pem
```

使用私钥生成钱包

```
DNA account import --pem --source ./pri
```

查看钱包的账户信息

```
DNA account list -v
```

## 2. 配置初始共识节点

编辑配置文件，修改共识节点信息为各节点的公钥和账户地址，并修改`SeedList`为服务器地址。然后重新编译。

## 3. 启动节点

启动节点

```
DNA --enable-consensus
```

程序会从`./cert.pem`中读取证书。
