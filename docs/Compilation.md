## 🏗️ Compilación


## Uniplataforma


Compilación para el sistema base.
  ```bash
  go build -o sniim-cli ./cmd/cli/main.go
  ```

## Multiplataforma

Adicionalmente, gracias a que el proyecto esta desarrollado con Go lang podemos compilar y generar un binario para cualquier plataforma y arquitectura, sin embargo para crear paquetes de distribución es preferible usar go releaser como se describe en la sección de [Releases](./Release.md).


- Compilación Linux 🐧
  ```bash
  GOOS=linux GOARCH=amd64 go build -o sniim-cli-linux-amd64 ./cmd/cli/main.go
  ```

- Compilación MacOS (amd64) 🍎
  ```bash
  GOOS=darwin GOARCH=amd64 go build -o sniim-cli-darwin-amd64 ./cmd/cli/main.go
  ```

- Compilación MacOS (arm64) 🍎
  ```bash
  GOOS=darwin GOARCH=arm64 go build -o sniim-cli-darwin-arm64 ./cmd/cli/main.go
  ```

- Compilación Windows 🟦
  ```bash
  GOOS=windows GOARCH=amd64 go build -o sniim-cli-windows-amd64 ./cmd/cli/main.go
  ```