# 下载
- 下载realise下对应系统的安装包
- 或者go环境下进行源码编译`./build.sh`生成三个系统的二进制程序
```shell script
#生成当前平台程序
go build wormhole
#生成mac,linux,windows3个平台程序(命名已区分)
./build.sh
```
# 使用
程序默认使用1808端口,可在config.yaml中修改,默认接收文件放在程序下的accept目录下,可自行修改
- 查看帮助
```shell script
$ wormhole
使用示例: 
1、启动本地服务：wormhole server 
2、查看在线设备：wormhole ls
3、发送消息：wormhole send 0 hello
4、发送文件：wormhole send ./wormhole

Usage:
  wormhole [flags]
  wormhole [command]

Available Commands:
  env         展示环境参数
  help        Help about any command
  list        列出所有设备
  net         添加指定IP或IP所在网段(IP最后为255)
  restart     重启服务
  send        发送文本消息或文件
  server      启动wormhole服务
  set         设置环境参数
  start       后台方式启动wormhole服务
  stop        停止服务

Flags:
  -h, --help   help for wormhole

Use "wormhole [command] --help" for more information about a command.
```

- 接受方需先启动服务:
```shell script
wormhole server #前台进程
wormhole start #后台进程
```
- 发送方可指定ip和端口直接发送,不指定flag时优先匹配发送信息为文件路径
```shell script
wormhole send 192.168.1.1:1808 ./file
wormhole send 192.168.1.1:1808 "this is a message"
```
- 发送方先启动服务后可以有更多发送方式
先查看主机列表信息:
```shell script
$ wormhole ls
+-------+---------------+------+-----------+--------+---------------------+
| INDEX |     HOST      | PORT |   NAME    | STATE  |   LASTACTIVETIME    |
+-------+---------------+------+-----------+--------+---------------------+
|     0 | 192.168.12.49 | 1808 | bishisimo | online | 2020-10-26 13:22:18 |
+-------+---------------+------+-----------+--------+---------------------+
```
```shell script
#通过索引发送,0号主机永远为本机
wormhole send 0 ./file
#通过用户名发送
wormhole send bishisimo ./file
#通过主机IP发送
wormhole send 192.168.1.1 ./file
```

# TODO
- [x] 维护README
- [x] 发送字符串消息
- [x] 发送文件
- [x] 发送文件夹
- [x] 不启动Server进行快捷发送
- [x] 使用命令进行配置变量打印与设置
- [x] 使用start(`wormhole start`)命令进行后台运行并通过stop或restart实现控制
- [ ] 支持中转服务器实现跨局域网通信(`wormhole fort`)