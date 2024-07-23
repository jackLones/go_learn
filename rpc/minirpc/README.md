## 手搓实现RPC

### 网络传输数据格式
* 两端要约定好数据包的格式
* 成熟的RPC框架会有自定义传输协议，这里网络传输格式定义如下，前面是固定长度消息头，后面是变长消息体
  ![](../../images/data_format.png)
### 自定义数据格式的读写
* 代码实现：[data_format.go](data_format.go)
* 相应的测试代码：[data_format_test.go](data_format_test.go)

### 编码解码
* 编码解码实现：[codec.go](codec.go)

### 实现RPC服务端
* 服务端接收到的数据需要包括什么？
  * 调用的函数名、参数列表，还有一个返回值error类型
* 服务端需要解决的问题是什么？
  * Map维护客户端传来调用函数，服务端知道去调谁
* 服务端的核心功能有哪些？
  * 维护函数map
  * 客户端传来的东西进行解析
  * 函数的返回值打包，传给客户端
* 代码实现：[server.go](server.go)

### 实现RPC客户端
* 客户端只有函数原型，使用reflect.MakeFunc() 可以完成原型到函数的调用
* reflect.MakeFunc()是Client从函数原型到网络调用的关键
* 代码实现：[client.go](client.go)

### 整体测试：实现RPC通信测试
* 给服务端注册一个查询用户的方法，客户端使用RPC方式调用
* 测试代码：[rpc_test.go](rpc_test.go)