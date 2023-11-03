![](https://img.shields.io/badge/version-v1.0.0-green.svg) &nbsp; ![](https://img.shields.io/badge/version-go1.21-green.svg) &nbsp;  ![](https://img.shields.io/badge/builder-success-green.svg) &nbsp;

> Mysql自动装配，只需简单配置即可使用

## 一、🚀🚀安装

### 1. Go get

```shell
go get github.com/archine/mysql-starter@v1.0.0
```

### 2、Go Mod

```shell
# go.mod文件加入下面的一条
github.com/archine/mysql-starter v1.0.6
# 命令行在该项目目录下执行
go mod tidy
```

## 二、使用文档

### 1、配置信息

| 配置                     | 描述                                             |
|------------------------|------------------------------------------------|
| mysql => log_level     | 日志级别，支持 info、error, 默认error                    |
| mysql => url           | 数据库连接地址，可直接输入完整的url，这时 database 不填即视为当前为完整的url |
| mysql => username      | 账号                                             |
| mysql => password      | 密码                                             |
| mysql => database      | 数据库，不填的情况下说明 url 为一个完整的                        |
| mysql => max_idle      | 连接池最大空闲连接，默认 10                                |
| mysql => max_connect   | 连接池最大连接数，默认 50                                 |
| mysql => max_idle_time | 连接空闲时间，超过即销毁，默认 30s                            |

示例:

```yaml
mysql:
  log_level: debug
  url: 127.0.0.1:3306
  username: root
  password: root
  database: go
  max_idle: 20
  max_connect: 50
  max_idle_time: 30s
```

### 2、项目使用

在任意 Bean 中直接注入即可

```go
package model

import (
	"github.com/archine/mysql-starter"
	"github.com/archine/ioc"
)


type User struct {
    Id int
	Username string
}

func (u *User) TableName() string {
	return "user"
}

type UserMapper struct {
	// 直接注入
	*starter.M
}

func (u *UserMapper) CreateBean() ioc.Bean {
	return &UserMapper{}
}
```