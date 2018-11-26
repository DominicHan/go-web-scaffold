# go-web-scaffold
## go语言web开发脚手架

# 项目初始化
## 1. 在$GOPATH/SRC下 git clone https://github.com/DominicHan/go-web-scaffold.git
## 2. 保证$GOPATH/bin 在环境变量$PATH下
## 3. 执行 go get -u -v github.com/kardianos/govendor
## 4. cd go-web-scaffold/ ,ls查看, 保证有vendor目录
## 5. govendor sync -v 同步 项目依赖包
## 6. 指定版本,添加第三方库 举例:govendor fetch -v github.com/gin-gonic/gin@v1.3.0 不指定版本,下载安装最新的就把包括@在内和后边去掉

2018.11.26

·项目按MVC模式进行改造

·添加了项目整体配置文件

·添加了mysql、rabbitmq、redis的配置

·添加了一些demo
