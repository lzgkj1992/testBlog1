# golang实现的博客程序

> GoMD是一款基于golang的beego框架开发的简洁markdown内容管理程序

+ [前台演示](http://xblogs.cn:8081)
+ [后台演示](http://xblogs.cn:8081/admin)
+ 账号密码 admin/admin(请勿修改账户信息)

## 如何使用

> 本应用基于golang语言的[beego框架](https://beego.me/)开发,在确保安装了golang环境的条件下，执行如下命令

```bash
go get github.com/astaxie/beego
go get github.com/beego/bee
go get github.com/jmoiron/sqlx
go get github.com/mattn/go-sqlite3
go get gopkg.in/russross/blackfriday.v2
```

> 安装完以上库，克隆仓库到机器上

```shell
cd $GOPATH/src/
git clone https://gitee.com/xuthus5/GoMD.git
cd GoMD
```

## 关于数据库与配置文件

数据库使用sqlite3，无需配置，编译运行程序即可使用，项目运行起来后，访问 http://domain/admin 为后台地址 默认账号密码 admin/admin

> 编译运行，两种方法 建议使用方法二

```
#运行方法一
go run main.go
#运行方法二
bee run
```

## 内置两套主题

1. QuietV1 [1025](https://1025.me/)

![首页](http://dl.xuthus.cc/q-i.png)
![内容](http://dl.xuthus.cc/q-a.png)

2. fantasy [Seevil](https://github.com/Seevil/fantasy)

![首页](http://dl.xuthus.cc/f-i.png)
![内容](http://dl.xuthus.cc/f-a.png)

## 后台

![首页](http://dl.xuthus.cc/admin.png)
![内容](http://dl.xuthus.cc/admin-a.png)