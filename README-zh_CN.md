# Sovereign (太阴星君) 🐰

<div align="right">

[English](README.md) | [中文](README-zh_CN.md)

</div>

[![Go 版本](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![许可证](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Kratos](https://img.shields.io/badge/Kratos-v2-00ADD8?style=flat&logo=go)](https://github.com/go-kratos/kratos)

> 基于 Kratos 框架构建的分布式服务治理平台，提供统一的服务治理能力。

## 📖 项目介绍

Sovereign (太阴星君)

## ✨ 核心特性


## 🚀 快速开始

### 前置要求

- Go 1.25+ (从源码构建)
- Docker & Docker Compose (容器化部署)
- MySQL 8.0+ (可选，用于数据库存储模式)
- etcd (可选，用于服务注册)

### 安装

#### 从源码安装

```bash
# 克隆仓库
git clone https://github.com/aide-family/sovereign.git
cd sovereign

# 初始化环境
make init

# 构建二进制文件
make build

# 运行服务
./bin/sovereign run all
```

#### 使用 Docker

```bash
# 构建 Docker 镜像
docker build -t sovereign:latest .

# 运行容器
docker run -d \
  --name sovereign \
  -p 8080:8080 \
  -p 9090:9090 \
  -v $(pwd)/config:/moon/config \
  sovereign:latest
```

## 📦 镜像构建

```bash
docker build -t sovereign-local:latest .
```

## 📦 部署

### Docker 部署

详细说明请参考 [Docker 部署文档](deploy/server/docker/README-docker.md)。

```bash
docker run -d \
  --name sovereign \
  -p 8080:8080 \
  -p 9090:9090 \
  -v $(pwd)/config:/moon/config \
  --restart=always \
  sovereign-local:latest run all
```

### docker-compose 部署

详细说明请参考 [Docker Compose 文档](deploy/server/docker/README-docker-compose.md)。

```bash
docker build -t sovereign-local:latest .
docker-compose -f deploy/server/docker/docker-compose.yml up -d
```

### Kubernetes 部署

详细说明请参考 [Kubernetes 部署文档](deploy/server/k8s/README.md)。

#### 快速部署

```bash
# 创建命名空间（如果不存在）
kubectl create namespace moon --dry-run=client -o yaml | kubectl apply -f -

# 部署 Sovereign 服务
cd deploy/server/k8s
kubectl apply -f sovereign.yaml
```

## 🤝 贡献指南

我们欢迎贡献！提交 PR 前请先阅读贡献指南。

### Pull Request 流程

1. **Fork 仓库**并从 `main` 分支创建你的分支
2. **创建 Issue** 讨论你的更改（如果是重大更改）
3. **进行更改**，遵循我们的代码风格指南
4. **添加测试**（新功能或 bug 修复）
5. **更新文档**（如需要）
6. **确保所有测试通过** (`make test`)
7. **提交 Pull Request**，附上清晰的描述

#### PR 标题格式

```
<type>(<scope>): <subject>

<body>

<footer>
```

**类型：**
- `feat`: 新功能
- `fix`: Bug 修复
- `docs`: 文档更改
- `style`: 代码风格更改（格式化等）
- `refactor`: 代码重构
- `test`: 测试添加或更改
- `chore`: 构建过程或辅助工具更改

**示例：**
```
feat(message): 添加邮件模板支持

添加对邮件模板的支持，支持动态变量替换。
模板可以在配置文件中定义，发送邮件时通过名称引用。

Closes #123
```

#### PR 检查清单

- [ ] 代码遵循项目的风格指南
- [ ] 已完成自我审查
- [ ] 为复杂代码添加了注释
- [ ] 已更新文档
- [ ] 已添加/更新测试
- [ ] 所有测试通过
- [ ] 未引入新的警告
- [ ] 更改向后兼容（或提供了迁移指南）

### Issue 报告

报告问题时，请包含：

1. **问题类型**：Bug、功能请求、问题等
2. **描述**：问题的清晰描述
3. **复现步骤**：对于 bug，提供复现步骤
4. **预期行为**：你期望发生什么
5. **实际行为**：实际发生了什么
6. **环境**：操作系统、Go 版本、Sovereign 版本
7. **配置**：相关配置（已脱敏）
8. **日志**：相关日志输出
9. **截图**：如适用


## 📄 许可证

本项目采用 MIT 许可证 - 详情请参阅 [LICENSE](LICENSE) 文件。

## 🙏 致谢

- [Kratos](https://github.com/go-kratos/kratos) - 微服务框架
- [Cobra](https://github.com/spf13/cobra) - Go 命令行框架

## 📞 联系方式

- **仓库**: https://github.com/aide-family/sovereign
- **Issues**: https://github.com/aide-family/sovereign/issues
- **邮箱**: aidecloud@163.com
- **飞书**:

  | ![](./docs/imgs/aide.png) | ![](./docs/imgs/enterprise.png) |
  | ------------------------- | ---- |

---

由 [Aide Family](https://github.com/aide-family) 用 ❤️ 制作
