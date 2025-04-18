# Go DDD Template Project

`scrape-hub` 是一個中央監控平台，用於集中顯示部署在 AWS Lambda 上的定時網路爬蟲任務狀態。每個爬蟲任務獨立執行並在完成後將執行結果（如成功/失敗與 log）回報至 `scrape-hub`，以便在 Web UI 中統一查詢與監控。

資料儲存採用 MongoDB，提供輕量且可擴展的事件儲存與查詢機制。

This is a template project implementing Clean Architecture and Domain-Driven Design (DDD) principles in Go. Use this template as a starting point for your domain-driven microservices.

## Project Structure

```
.
├── cmd/            # Application entrypoints
├── configs/        # Configuration files and structures
├── internal/       # Private application code
│   ├── domain/     # Domain layer: entities, value objects, domain services
│   ├── usecase/    # Use case layer: commands, queries, behaviors
│   ├── delivery/   # Interface adapters (HTTP/gRPC)
│   ├── infra/      # Infrastructure implementations
│   └── shared/     # Shared utilities
├── pkg/            # Public libraries
└── tests/          # Test suites
```

## Getting Started

1. Click "Use this template" on GitHub to create a new project
2. Clone your new repository
3. Update the module name in `go.mod`
4. Start implementing your domain model in `internal/domain`

## Architecture Overview

This template follows Clean Architecture and DDD principles:

- **Domain Layer**: Core business logic and rules, including:
  - Entities and Aggregates
  - Value Objects
  - Domain Services
  - Repository Interfaces
  
- **Use Case Layer**: Application flows and coordination:
  - Commands (write operations)
  - Queries (read operations)
  - Behaviors (complex workflows)
  - Event Handlers

- **Interface Adapters**: Multiple delivery mechanisms:
  - HTTP REST APIs
  - gRPC Services
  - Message Consumers

- **Infrastructure**: External concerns:
  - Database Implementations
  - External Service Clients
  - Message Brokers

## Key Design Principles

1. Dependencies flow inward (domain at the center)
2. Domain layer has no external dependencies
3. Use interfaces for infrastructure concerns
4. Separation of commands and queries (CQRS)
5. Domain events for cross-boundary communication

## Documentation

Each directory contains its own README.md with specific guidance for that component.

## Testing

- Unit tests alongside the code
- Integration tests in /tests
- E2E tests in /tests/e2e
