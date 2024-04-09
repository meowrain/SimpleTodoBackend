# ✨TodoBackend-TodoApp后端

<center>
<img src="docs/logo.png" width="200" height="200">
</center>


> TodoApp后端，使用Gin Framework框架实现

## 😊功能

- [x] Todo 增删改查
- [x] 用户增删改查
- [x] 用户登录(jwt鉴权)
- [x] 用户注册
- [x] Todo 图片上传
- [x] Todo条数统计
- [x] 用户个人简介删改
## 开发中

### 开发进度

- [x] 项目初始化
- [x] 项目结构
- [x] 项目配置
- [x] 项目路由
- [x] 项目中间件
- [x] 项目模型
- [x] 项目控制器
- [x] 项目服务
- [x] 项目文档
- [x] 项目测试
- [x] 项目部署


# 直接使用源代码启动项目

## 克隆项目

```shell
git clone https://github.com/meowrain/TodoBackend.git
```

## 进入项目目录

```shell
cd TodoBackend
```

## 安装依赖

```shell
go get
```
## 启动项目

```shell
go run main.go
```


# 从源代码构建程序

首先请确保你的电脑安装了go语言环境还有GNU Make

## 为所有平台进行构建

```shell
make build-all
```

## 仅为windows平台构建

```shell
make build-windows
```

## 仅为macOS平台构建

```shell
make build-macos
```

## 仅为linux平台构建

```shell
make build-linux
```

### 运行程序

> 复制config.yaml到可执行文件目录下

运行
```shell
./SimpleTodoBackend-${platform}-${arch}
```
