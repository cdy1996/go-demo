
```shell script
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/grpc
# 编译proto文件
protoc.exe -I ./ helloworld.proto --go_out=plugins=grpc:.

```


[protoc与protoc-gen-go安装](https://blog.csdn.net/VBVSPER/article/details/91878373)
> 当Linux系统无法访问网络时：首先在github.com/golang/protobuf上下载protoc-gen-go和proto，（最好将其放在$GOPATH/src目录下）然后进入protoc-gen-go目录，执行go build、go install即可在$GOPATH/bin目录下发现这个工具。
 
