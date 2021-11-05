# Go环境安装
下载go
方法一：官网下载：https://golang.org/dl/
方法二：打开ubuntu输入：
wget https://dl.google.com/go/go1.15.3.linux-amd64.tar.gz

解压安装包：
sudo tar -C /usr/local -xzf go1.15.3.linux-amd64.tar.gz

建立软连接：
sudo ln -s /usr/local/go/bin/* /usr/bin/

sudo vim ~/.bashrc

进入编辑界面后Shift+G跳转至尾行，按o新插入一行，输入：
export GOPATH="$HOME/go"
export PATH="$PATH:/usr/local/go/bin:$GOPATH/bin"

Esc退出编辑，输入以下代码退出文件保存并生效：
:wq
source ~/.bashrc

查看安装是否生效：
go version


# GOLand 新建项目 
- 1.文件 setting - Go 模块 - 启用GO模块
- 2.命令行输入 go mod init github.com/HuCuiGang/bank
