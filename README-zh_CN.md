# Sovereign (太阴星君)

<div align="right">

[English](README.md) | [中文](README-zh_CN.md)

</div>

[![许可证](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go 版本](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Kratos](https://img.shields.io/badge/Kratos-v2.9.2-00ADD8?style=flat&logo=go)](https://github.com/go-kratos/kratos)
[![Cobra](https://img.shields.io/badge/Cobra-v1.10.2-00ADD8?style=flat&logo=go)](https://github.com/spf13/cobra)

## 📖 项目介绍

Sovereign (太阴星君) 是作为 moon 体系通用的服务模板项目

## 🚀 快速开始
```bash
make init
make build
```

### 运行二进制文件

- 帮助

```bash
./bin/sovereign -h
```

- 版本

```bash
./bin/sovereign version
```

- 运行所有服务

```bash
./bin/sovereign run all -h
```

- 运行 gRPC 服务

```bash
./bin/sovereign run grpc -h
```

- 运行 HTTP 服务

```bash
./bin/sovereign run http -h
```

## 开发

```bash
make init
make all
```

### 运行应用程序

- run all

```bash
go run . run all
```

- run grpc

```bash
go run . run grpc
```

- run http

```bash
go run . run http
```

## 致谢

- [kratos](https://github.com/go-kratos/kratos)
- [cobra](https://github.com/spf13/cobra)