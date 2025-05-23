# Golang SDK For FISCO BCOS

适配FISCO-BCOS v3 / [适配FISCO-BCOS v2](https://github.com/FISCO-BCOS/go-sdk/tree/master-FISCO-BCOS-v2)

[![CodeFactor](https://www.codefactor.io/repository/github/fisco-bcos/go-sdk/badge)](https://www.codefactor.io/repository/github/fisco-bcos/go-sdk)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/afbb696df3a8436a9e446d39251b2158)](https://www.codacy.com/gh/FISCO-BCOS/go-sdk?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=FISCO-BCOS/go-sdk&amp;utm_campaign=Badge_Grade)
[![codecov](https://codecov.io/gh/FISCO-BCOS/go-sdk/branch/master/graph/badge.svg)](https://codecov.io/gh/FISCO-BCOS/go-sdk)
[![Code Lines](https://tokei.rs/b1/github/FISCO-BCOS/go-sdk?category=code)](https://github.com/FISCO-BCOS/go-sdk)
[![version](https://img.shields.io/github/tag/FISCO-BCOS/go-sdk.svg)](https://github.com/FISCO-BCOS/go-sdk/releases/latest)

![FISCO-BCOS Go-SDK GitHub Actions](https://github.com/FISCO-BCOS/go-sdk/workflows/FISCO-BCOS%20Go-SDK%20GitHub%20Actions/badge.svg)
[![Build Status](https://travis-ci.org/FISCO-BCOS/go-sdk.svg?branch=master)](https://travis-ci.org/FISCO-BCOS/go-sdk)

____

FISCO BCOS Go语言版本的SDK，主要实现的功能有：

- [FISCO BCOS 3.0 JSON-RPC服务](https://fisco-bcos-doc.readthedocs.io/zh_CN/latest/docs/develop/api.html)
- `Solidity`合约编译为Go文件
- 部署、查询、写入智能合约
- 控制台

`go-sdk`的使用可以当做是一个`package`进行使用，亦可对项目代码进行编译，直接使用**控制台**通过配置文件来进行访问FISCO BCOS，3.0版本的go sdk使用cgo依赖bcos-c-sdk以支持国密等特性，请注意先按照文档下载bcos-c-sdk的动态库。

## 环境准备

master分支的go-sdk对应FISCO-BCOS v3版本，如果使用的是FISCO-BCOS v2版本，请切换到[FISCO-BCOS v2分支](https://github.com/FISCO-BCOS/go-sdk/tree/master-FISCO-BCOS-v2)，[对应文档](https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/sdk/go_sdk/index.html)

- [Golang](https://golang.org/), 版本需不低于`1.21`，本项目采用`go module`进行包管理。具体可查阅[Using Go Modules](https://blog.golang.org/using-go-modules)
- [FISCO BCOS 3.2.0+](https://fisco-bcos-doc.readthedocs.io/zh_CN/latest/index.html), 可参考[安装搭建](https://fisco-bcos-doc.readthedocs.io/zh_CN/latest/docs/quick_start/air_installation.html)
- Solidity编译器，默认0.8.11版本

## 配置结构体说明

```golang
type Config struct {
    TLSCaFile       string
    TLSKeyFile      string
    TLSCertFile     string
    TLSSmEnKeyFile  string
    TLSSmEnCertFile string
    IsSMCrypto      bool
    PrivateKey      []byte
    GroupID         string
    Host            string
    Port            int
    DisableSsl      bool
}
```

- `TLSCaFile/TLSKeyFile/TLSCertFile`，建立TLS链接时需要用到的SDK端证书文件路径，如果是国密，其加密私钥和证书使用`TLSSmEnKeyFile/TLSSmEnCertFile`
- IsSMCrypto：节点使用的签名和TLS算法，true表示使用国密，false表示使用RSA+ECDSA。
- PrivateKey：节点签名交易时所使用的私钥，支持国密和非国密。(pem文件可使用LoadECPrivateKeyFromPEM方法解析)
  请使用[get_account.sh](https://github.com/FISCO-BCOS/console/blob/master/tools/get_account.sh)和[get_gm_account.sh](https://github.com/FISCO-BCOS/console/blob/master/tools/get_gm_account.sh)脚本生成。使用方式[参考这里](https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/manual/account.html)。
  如果想使用Go-SDK代码生成，请[参考这里](doc/README.md#外部账户)。
- GroupID：账本的`GroupID`
- Host：节点IP
- Port：节点RPC端口
- DisableSsl：使用TLS加密时为`false`，不使用TLS加密时为`true`，此配置项需与节点config.ini中的`rpc.disable_ssl`保持一致

## 控制台使用

1. 搭建FISCO BCOS 3.2以上版本节点，请[参考这里](https://fisco-bcos-doc.readthedocs.io/zh_CN/latest/docs/quick_start/air_installation.html)。
1. 请拷贝对应的SDK证书到conf文件夹，证书名为`ca.crt/sdk.key/sdk.crt`，国密时证书名为`sm_ca.crt/sm_sdk.key/sm_sdk.crt/sm_ensdk.key/sm_ensdk.crt`。
1. go-sdk需要依赖csdk的动态库，[下载地址](https://github.com/FISCO-BCOS/bcos-c-sdk/releases/tag/v3.6.0)，将动态库放在`/usr/local/lib`目录下。在其他机器使用时也需要通过`export LD_LIBRARY_PATH=${PWD}/lib`设置动态库的搜索路径，其中`${PWD}/lib`需替换为bcos-c-sdk的动态库所在文件夹。如果编译后在其他机器运行，也可以在编译时使用`-ldflags`指定动态库搜索路径，如`go build -ldflags="-r ${PWD}/lib"`。

    ```bash
    # 下面的脚本帮助用户下载bcos-c-sdk的动态库到/usr/local/lib目录下
    ./tools/download_csdk_lib.sh
    ```

1. ~~go-sdk需要使用cgo，需要设置环境变量`export GODEBUG=cgocheck=0`~~。
1. 开启CGO支持`go env -w CGO_ENABLED=1`。
1. 最后，编译控制台程序:

```bash
git clone https://github.com/FISCO-BCOS/go-sdk.git
cd go-sdk
cd v3
go mod tidy
go build -ldflags="-r /usr/local/lib" -o console ./cmd/console.go
./console help
```

## Package功能使用

以下的示例是通过`import`的方式来使用`go-sdk`，如引入RPC控制台库:

```go
import "github.com/FISCO-BCOS/go-sdk/v3/client"
```

### Solidity合约编译为Go文件

在利用SDK进行项目开发时，对智能合约进行操作时需要将Solidity智能合约利用go-sdk的`abigen`工具转换为`Go`文件代码。整体上主要包含了五个流程：

- 准备需要编译的智能合约
- 配置好相应版本的`solc`编译器
- 构建go-sdk的合约编译工具`abigen`
- 编译生成go文件
- 使用生成的go文件进行合约调用

下面的内容作为一个示例进行使用介绍。

1.提供一份简单的样例智能合约　v3/examples/file_box_v2/FileBoxV2.so

2.安装对应版本的[`solc`编译器](https://github.com/ethereum/solidity/releases/tag/v0.8.11)，目前FISCO BCOS默认的`solc`编译器版本为0.8.11。

```bash
# 如果是国密则添加-g选项
bash tools/download_solc.sh -v 0.8.11 -g
```

3.构建`go-sdk`的代码生成工具`abigen`

```bash
# 下面指令在go-sdk目录下操作，编译生成abigen工具
cd v3
go build ./cmd/abigen
cd ..
cp v3/abigen .
```

4.编译生成go文件，先利用`solc`将合约文件生成`abi`和`bin`文件，以`FileBoxV2.sol`为例：

```bash
# 国密请使用 ./solc-0.8.11-gm --bin --abi -o ./hello ./hello/HelloWorld.sol
./solc-0.8.11-gm --bin --abi -o v3/examples/file_box_v2 v3/examples/file_box_v2/FileBoxV2.sol
```

在MacOS下运行`./solc-0.8.11`时如果出现找不到`libz3.dylib`的错误,例如:

```bash
dyld[42564]: Library not loaded: /opt/homebrew/opt/z3/lib/libz3.dylib
  Referenced from: <08BAD135-54EC-3430-A170-26E7B4A5BA96> xxxxxx/.fisco/solc/solc-0.8.11
  Reason: tried xxxxxx
```

可尝试以下命令安装`libz3`

```bash
brew installz3
```

`FileBoxV2.sol`目录下会生成`FileBoxV2.bin`和`FileBoxV2.abi`。此时利用`abigen`工具将`FileBoxV2.bin`和`FileBoxV2.abi`转换成`FileBoxV2.go`：

```bash
# 国密请使用 ./abigen --bin ./hello/HelloWorld.bin --abi ./hello/HelloWorld.abi --pkg hello --type HelloWorld --out ./hello/HelloWorld.go --smcrypto=true
# 注意：国密模式，请使用国密solc编译得到bin
./abigen --bin v3/examples/file_box_v2/FileBoxV2.bin --abi v3/examples/file_box_v2/FileBoxV2.abi --pkg file_box_v2 --type FileBox --out v3/examples/file_box_v2/FileBoxV2.go
```

最后v3/examples/file_box_v2目录下面存在以下文件：

```bash
FileBoxV2.abi  FileBoxV2.bin  FileBoxV2.go  FileBoxV2.sol
```

5.调用生成的`FileBoxV2.go`文件进行合约调用

至此，合约已成功转换为go文件，用户可根据此文件在项目中利用SDK进行合约操作。具体的使用可参阅下一节。

### 部署智能合约

下面的例子先部署合约，在部署过程中设置的`FileBoxV2.sol`合约中有一个公开的名为`version`的全局变量，这种公开的成员将自动创建`getter`函数，然后调用`Version()`来获取version的值。

写入智能合约需要我们用私钥来对交易事务进行签名，我们创建的智能合约有一个名为`Set`的方法，它接受`string`类型的参数，然后将其设置为`value`，并且将`version`加1。

新建一个go工程，目录结构如下

```bash
hello
  ｜—— FileBoxV2.go
go.mod
go.sum
hello_main.go
```

`hello_main.go`代码如下

```go
package main

import (
 "context"
 "encoding/hex"
 "fmt"
 "log"

 "example/go-sdk/hello"

 "github.com/FISCO-BCOS/go-sdk/v3/client"
 "github.com/FISCO-BCOS/go-sdk/v3/types"
)

func main() {
 privateKey, _ := hex.DecodeString("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")
 config := &client.Config{IsSMCrypto: false, GroupID: "group0",
  PrivateKey: privateKey, Host: "127.0.0.1", Port: 20200, TLSCaFile: "./ca.crt", TLSKeyFile: "./sdk.key", TLSCertFile: "./sdk.crt"}
 client, err := client.DialContext(context.Background(), config)
 if err != nil {
  log.Fatal(err)
 }
 input := "HelloWorld deployment 1.0"
 fmt.Println("=================DeployHelloWorld===============")
 address, receipt, instance, err := hello.DeployHelloWorld(client.GetTransactOpts(), client, input)
 if err != nil {
  log.Fatal(err)
 }
 fmt.Println("contract address: ", address.Hex()) // the address should be saved, will use in next example
 fmt.Println("transaction hash: ", receipt.TransactionHash)

 // load the contract
 // contractAddress := common.HexToAddress("contract address in hex String")
 // instance, err := hello.NewHelloWorld(contractAddress, client)
 // if err != nil {
 //     log.Fatal(err)
 // }

 fmt.Println("================================")
 helloSession := &hello.HelloWorldSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

 version, err := helloSession.Version()
 if err != nil {
  log.Fatal(err)
 }

 fmt.Println("version :", version) // "HelloWorld deployment 1.0"

 ret, err := helloSession.Get()
 if err != nil {
  fmt.Printf("hello.Get() failed: %v", err)
  return
 }
 done := make(chan bool)
 currentBlock, err := client.GetBlockNumber(context.Background())
 if err != nil {
  fmt.Printf("GetBlockNumber() failed: %v", err)
  return
 }
 _, err = helloSession.WatchAllSetValue(&currentBlock, func(ret int, logs []types.Log) {
  fmt.Printf("WatchAllSetValue receive statud: %d, logs: %v\n", ret, logs)
  setValue, err := helloSession.ParseSetValue(logs[0])
  if err != nil {
   fmt.Printf("hello.WatchAllSetValue() failed: %v", err)
   panic("WatchAllSetValue hello.WatchAllSetValue() failed")
  }
  fmt.Printf("receive setValue: %+v\n", *setValue)
  done <- true
 })
 if err != nil {
  fmt.Printf("hello.WatchAllSetValue() failed: %v", err)
  return
 }
 fmt.Printf("Get: %s\n", ret)
 fmt.Println("================================")

 oldValue, _, receipt, err := helloSession.Set("hello fisco")
 fmt.Println("old value is: ", oldValue)
 if err != nil {
  log.Fatal(err)
 }

 fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())

 ret, err = helloSession.Get()
 if err != nil {
  fmt.Printf("hello.Get() failed: %v", err)
  return
 }
 fmt.Printf("Get: %s\n", ret)
 <-done
}

```

在项目目录下运行`hello_world.go`

```bash
go run -ldflags="-r /usr/local/lib" main.go
```
