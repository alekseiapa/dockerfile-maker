# 🐳 Dockerfile Maker

Welcome to the **Dockerfile Maker** project! This project is a robust Go-based application and library for generating Dockerfiles using various input channels. Designed to help you efficiently create Dockerfile instructions, especially useful when Dockerfile's limited control logic becomes a barrier.

---

## 📚 Overview

`dockerfile-maker` is both a Go library and an executable tool that generates valid Dockerfiles by leveraging structured input data, primarily through YAML. It is ideal for projects that require dynamic or automated Dockerfile creation, such as CI/CD pipelines or configs-as-code setups.

![Go Report Card](https://goreportcard.com/badge/github.com/alekseiapa/dockerfile-maker) ![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)

---

## 🎯 Features

- **Dynamic Generation:** Convert structured input from YAML to Dockerfile.
- **Multi-Mode Usage:** Utilize as a standalone app or integrate as a library within other Go applications.
- **YAML Configuration:** Effortlessly manage Dockerfile instructions through YAML.
- **Multi-Stage Support:** Generate Dockerfiles that optimize image size and build efficiency.

---

## ⚙️ Installation

### As an Executable

For direct command-line usage, download the executable for your respective OS and make it executable:

#### MacOS
```bash
curl -o dfg -L https://github.com/alekseiapa/dockerfile-maker/releases/download/v1.0.0/dfg_v1.0.0_darwin_amd64
chmod +x dfg && sudo mv dfg /usr/local/bin
```

#### Linux
```bash
curl -o dfg -L https://github.com/alekseiapa/dockerfile-maker/releases/download/v1.0.0/dfg_v1.0.0_linux_amd64
chmod +x dfg && sudo mv dfg /usr/local/bin
```

#### Windows
```bash
curl -o dfg.exe -L https://github.com/alekseiapa/dockerfile-maker/releases/download/v1.0.0/dfg_v1.0.0_windows_amd64.exe
```

### As a Library

To include `dockerfile-maker` in your Go project:

```bash
go get -u github.com/alekseiapa/dockerfile-maker
```

---

## 🚀 Getting Started

### Using as an Executable

With the executable, generating Dockerfiles is straightforward:

```bash
dfg generate --input path/to/yaml --out Dockerfile
```

You can also specify a field in the YAML to use for Dockerfile instructions:

```bash
dfg generate --input path/to/yaml --target-field ".server.dockerfile" --out Dockerfile
```

### Using as a Library

Integrate Dockerfile generation within your Go application:

```go
package main

import (
    "os"
    dfg "github.com/alekseiapa/dockerfile-maker"
)

func main() {
    data := &dfg.DockerfileData{
        Stages: []dfg.Stage{
            []dfg.Instruction{
                dfg.From{Image: "alpine:latest"},
                // More instructions...
            },
        },
    }
    tmpl := dfg.NewDockerfileTemplate(data)
    err := tmpl.Render(os.Stdout)
    if err != nil {
        panic(err.Error())
    }
}
```

---

## 📘 Documentation

This project provides detailed documentation on how to leverage the `dockerfile-maker` capabilities to structure and transform Dockerfile instructions:

- **Classes and Interfaces:** Utilize `Instruction`, `DockerfileData`, `Stage` for seamless process interaction.
- **YAML Integration:** Enhance configurations using YAML, ensuring structure and readability.

For comprehensive examples and usage guides, explore our [test cases](template_test.go) and the [source code](template.go).

---

## 🛠️ Development Setup

### Preparation

1. **Code Formatting:**
   - Use `make format` to format code automatically.
2. **Linting:**
   - Run `make lint` to identify and fix code quality issues.
3. **Testing:**
   - Execute tests using `make test` to ensure functionality.
4. **Benchmarking:**
   - Use `make benchmark` to assess code performance.

---

## 📜 Credits & Acknowledgements

This project, **`dfg - Dockerfile Generator`**, was inspired by the original project [dockerfile-generator](https://github.com/ozankasikci/dockerfile-generator) created by [Ozan Kasikci](https://github.com/ozankasikci).

All rights and credits for the original idea, design, and concept belong to the original author. I used this project solely for **learning and study purposes** to improve my Go programming skills and better understand Dockerfile generation techniques.

Thank you to the original author for creating such a great resource!
