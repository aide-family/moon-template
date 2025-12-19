# Sovereign (太阴星君) 🐰

<div align="right">

[English](README.md) | [中文](README-zh_CN.md)

</div>

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Kratos](https://img.shields.io/badge/Kratos-v2-00ADD8?style=flat&logo=go)](https://github.com/go-kratos/kratos)

> A distributed messaging platform built on the Kratos framework, providing unified message delivery and management capabilities.

## 📖 Introduction

Sovereign (太阴星君)

## ✨ Features


## 🚀 Quick Start

### Prerequisites

- Go 1.25+ (for building from source)
- Docker & Docker Compose (for containerized deployment)
- MySQL 8.0+ (optional, for database storage mode)
- etcd (optional, for service registry)

### Installation

#### From Source

```bash
# Clone the repository
git clone https://github.com/aide-family/sovereign.git
cd sovereign

# Initialize the environment
make init

# Build the binary
make build

# Run the service
./bin/sovereign run all
```

#### Using Docker

```bash
# Build the Docker image
docker build -t sovereign:latest .

# Run the container
docker run -d \
  --name sovereign \
  -p 8080:8080 \
  -p 9090:9090 \
  -v $(pwd)/config:/moon/config \
  sovereign:latest
```

## 📦 Image Build

```bash
docker build -t sovereign-local:latest .
```

## 📦 Deployment

### Docker Deployment

See [Docker Deployment Documentation](deploy/server/docker/README-docker.md) for detailed instructions.

```bash
docker run -d \
  --name sovereign \
  -p 8080:8080 \
  -p 9090:9090 \
  -v $(pwd)/config:/moon/config \
  --restart=always \
  sovereign-local:latest run all
```

### docker-compose Deployment

See [Docker Compose Documentation](deploy/server/docker/README-docker-compose.md) for detailed instructions.

```bash
docker build -t sovereign-local:latest .
docker-compose -f deploy/server/docker/docker-compose.yml up -d
```

### Kubernetes Deployment

See [Kubernetes Deployment Documentation](deploy/server/k8s/README.md) for detailed instructions.

#### Quick Deploy

```bash
# Create namespace (if not exists)
kubectl create namespace moon --dry-run=client -o yaml | kubectl apply -f -

# Deploy Sovereign service
cd deploy/server/k8s
kubectl apply -f sovereign.yaml
```


## 🤝 Contributing

We welcome contributions! Please read our contributing guidelines before submitting PRs.

### Pull Request Process

1. **Fork the repository** and create your branch from `main`
2. **Create an issue** to discuss your changes (if it's a significant change)
3. **Make your changes** following our code style guidelines
4. **Add tests** for new features or bug fixes
5. **Update documentation** as needed
6. **Ensure all tests pass** (`make test`)
7. **Submit a Pull Request** with a clear description

#### PR Title Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Test additions or changes
- `chore`: Build process or auxiliary tool changes

**Example:**
```
feat(service): add service治理能力

Add support for service治理能力.

Closes #123
```

#### PR Checklist

- [ ] Code follows the project's style guidelines
- [ ] Self-review completed
- [ ] Comments added for complex code
- [ ] Documentation updated
- [ ] Tests added/updated
- [ ] All tests pass
- [ ] No new warnings introduced
- [ ] Changes are backward compatible (or migration guide provided)

### Issue Reporting

When reporting issues, please include:

1. **Issue Type**: Bug, Feature Request, Question, etc.
2. **Description**: Clear description of the issue
3. **Steps to Reproduce**: For bugs, provide steps to reproduce
4. **Expected Behavior**: What you expected to happen
5. **Actual Behavior**: What actually happened
6. **Environment**: OS, Go version, Sovereign version
7. **Configuration**: Relevant configuration (sanitized)
8. **Logs**: Relevant log output
9. **Screenshots**: If applicable

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Kratos](https://github.com/go-kratos/kratos) - A microservice-oriented framework
- [Cobra](https://github.com/spf13/cobra) - A CLI framework for Go

## 📞 Contact

- **Repository**: https://github.com/aide-family/sovereign
- **Issues**: https://github.com/aide-family/sovereign/issues
- **Email**: aidecloud@163.com
- **Feishu**: 

  | ![](./docs/imgs/aide.png) | ![](./docs/imgs/enterprise.png) |
  | ------------------------- | ---- |

  

---

Made with ❤️ by [Aide Family](https://github.com/aide-family)
