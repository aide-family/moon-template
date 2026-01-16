# Sovereign (太阴星君)

<div align="right">

[English](README.md) | [中文](README-zh_CN.md)

</div>

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Kratos](https://img.shields.io/badge/Kratos-v2-00ADD8?style=flat&logo=go)](https://github.com/go-kratos/kratos)

## 📖 Introduction

Sovereign (太阴星君) is a universal service template project for the Moon platform.

## Quick Start

```bash
make init
make build
```

### run the binary

- help

```bash
./bin/sovereign -h
```

- version

```bash
./bin/sovereign version
```

- run all

```bash
./bin/sovereign run all -h
```

- run grpc

```bash
./bin/sovereign run grpc -h
```

- run http

```bash
./bin/sovereign run http -h
```

## Development

```bash
make init
make all
```

### run the application

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

## Acknowledgments

- [kratos](https://github.com/go-kratos/kratos)
- [cobra](https://github.com/spf13/cobra)