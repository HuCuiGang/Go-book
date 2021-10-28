# go 语言的优势与劣势
## 优势
- 1.脚本化的语法
- 2.静态类型+编译型 ，程序运行速度有保障 。 相比动态类型+解释型 高很多
- 3.原生的支持并发编程 
降低开发、维护成本 ，程序可以更好的执行。

## 劣势
- 1.语法糖并没有Python和Ruby那么多
- 2.目前的程序运行速度还不及C
- 3.第三方函数库暂时不多（比如Java，python）

# GO 语言的安装与设置
## linux下的安装
- 1.方法1：
从http://golang.org/dl/下载最新版本的Go语言二进制档案包
- 2.方法2：
比如∶若要将版本为1.4.2的Go语言安装到计算架构为64位的Linux操作系统上，则需要下载名为go1.4.2.linux-amd64.tar.gz 的档案包。
使用tar命令将档案包解压到/usr/local目录中
- 验证安装结果
进入到/usr/local目录中查看是否存在一个名为go的目录
在命令行下进入到这个go目录，敲入bin/go version并回车，查看是否有如下图所示的Go语言版本信息打印出来。

## linux下的设置方法
- 有4个环境变量需要设置:GOROOT、GOPATH、GOBIN，以及PATH需要设置到某一个profile 文件中(~/.bash_profile或/etc/profile)
- export GOROOT=/usr/local/go
- export GOPATH=~/golib:~/goproject
- export GOBIN=~/gobin
- export PATH=$PATH: $GOROOT/bin:$GOBIN

- source <某个profile文件的绝对路径>

- 在命令行下的任意目录中敲入go version并回车，然后检查打印信息

## 工作区目录含义
- /home/hypermind/golib:
src/  源码文件
pkg/  归档问文件 .a 后缀
bin/  go可执行文件










