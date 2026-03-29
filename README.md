# ChaincodeInfo

在`/etc/hosts`中添加：
```
127.0.0.1  orderer.example.com
127.0.0.1  peer0.org1.example.com
127.0.0.1  peer1.org1.example.com
```
添加依赖：
```
cd ChaincodeInfo && go mod tidy
```
运行项目：
```
./deploy.sh
```
在`127.0.0.1:3000`进行访问
# ChaincodeInfo
