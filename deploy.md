# 联盟链demo部署说明

version 0.1

## 1. 准备节点钱包

通过fabric ca为节点生成密钥和证书，目录结构如下：

```
.
├── fabric-ca-client-config.yaml
└── msp
    ├── IssuerPublicKey
    ├── IssuerRevocationPublicKey
    ├── cacerts
    │   └── localhost-7054.pem
    ├── keystore
    │   └── 414081a4e557442b5e7a1bd6762a02cf884679ce6809629c343c212a02fbf159_sk
    ├── signcerts
    │   └── cert.pem
    └── user
```

使用`./msp/keystore`目录中的私钥文件生成钱包

```
ontology account import --pem --source ./msp/keystore/414081a4e557442b5e7a1bd6762a02cf884679ce6809629c343c212a02fbf159_sk
```

查看钱包的账户信息

```
ontology account list -v
```

## 2. 配置初始共识节点

编辑源码文件`common/config/config.go`，修改`MainNetConfig`中的节点信息为各节点的公钥和账户地址，并修改`SeedList`为服务器地址。然后重新编译。

## 3. 启动节点

在`msp`的同级目录启动节点

```
ontology --enable-consensus
```

程序会从`./msp/signcerts/`目录中读取证书。