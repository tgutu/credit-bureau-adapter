## Prerequisites

- Go >= 1.25 (If compiling on machine and not in Docker)
- Docker

---

## Setup

1. **Clone the repository:**
    ```sh
    git clone https://github.com/tgutu/credit-bureau-adapter.git
    cd credit-bureau-adapter
    ```

2. **Run the service in Docker**
    ```sh
    docker compose up
    ```

3. **Install dependencies:**
    ```sh
    go mod tidy
    ```

4. **Install build tools**
    ```sh
    brew install buf
    brew install golangci-lint
    ```

---

## Development Workflow

- **Branching:** Use feature branches (`feature/xyz`), bugfix branches (`bugfix/xyz`), or hotfix branches (`hotfix/xyz`).
- **Commits:** Write clear, concise commit messages.
- **Pull Requests:** Submit PRs for all changes. Ensure all checks pass before merging.

---

## Protobuf
- The service is defined using Google Protobuf
- The gRPC service definition is located in the `proto/` folder
- Ensure protobuf updates are sane: `buf lint`
- Generate new source after updates: `buf generate`
---

## Testing

- **Unit Tests:**  
  Run tests using:
  ```sh
  go test ./...
  ```

- **Postman:**
  Import the OpenAPI spec inside the `docs/` folder into Postman.
  ```sh
  curl --location 'http://localhost:8080/v1/credit/bureaus' # List credit bureaus
  ```

- **Integration Tests:**  
  Use local endpoint for integration tests.

---

## Code Style

- Follow Go conventions and use `gofmt` for formatting.
- Run lint checks before committing:
  ```sh
  golangci-lint run
  ```

---
