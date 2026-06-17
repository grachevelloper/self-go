# Application Layers Reference

This project follows layered architecture with dependencies pointing inward.

```text
cmd/api
  -> delivery/http
    -> usecase
      -> domain

infrastructure/postgres
  -> usecase contracts
  -> domain
```

Infrastructure and delivery are outer layers. Domain and usecase are inner layers.

## Layer Responsibilities

| Layer | Responsibility | May Know About | Must Not Know About |
|---|---|---|---|
| `cmd/api` | Application composition: config, database connection, repositories, usecases, handlers, HTTP server startup | Concrete implementations from all layers | Low-level business decisions hidden inside SQL or handlers |
| `internal/domain` | Core domain entities and domain rules | Domain concepts, standard library when needed | HTTP, JSON, SQL, Postgres, handlers, repositories, framework details |
| `internal/usecase/book` | Application operations for books: create, read, update, delete, list, search | `domain`, usecase input/output types, repository interfaces | HTTP, JSON tags, SQL queries, `*sql.DB`, Postgres-specific behavior |
| `internal/usecase/book.Repository` | Storage contract required by book usecases | `context`, `domain`, usecase-level input/output types | Concrete database implementation, HTTP DTOs, handlers |
| `internal/delivery/http/book` | HTTP transport: request parsing, validation mapping, response formatting, status codes | `net/http`, HTTP DTOs, usecase interfaces | SQL, Postgres, `*sql.DB`, repository implementation details |
| `internal/delivery/http/book/dto.go` | HTTP request and response shapes | JSON tags, HTTP field names, transport-specific optional fields | SQL columns, database behavior, repository contracts |
| `internal/infrastructure/postgres/book` | Postgres implementation of book persistence | `database/sql`, SQL queries, `domain`, usecase repository contracts | HTTP DTOs, handlers, JSON tags, transport-specific behavior |
| `internal/config` | Configuration loading | Environment variables, configuration formats | Business logic, HTTP request handling, SQL query behavior |

## Dependency Rules

| Rule | Correct | Incorrect |
|---|---|---|
| Domain is independent | `domain.Book` has business fields only | `domain.Book` has JSON or SQL tags because a transport needs them |
| Delivery calls usecase | HTTP handler depends on a book service/usecase interface | HTTP handler calls Postgres repository directly |
| Usecase owns business contracts | Usecase defines inputs needed for create/update operations | Usecase accepts HTTP DTOs with JSON tags |
| Repository implements usecase needs | Postgres repository satisfies the usecase repository interface | Usecase imports the Postgres repository package |
| Infrastructure stays transport-agnostic | Postgres repository accepts domain/usecase types | Postgres repository imports `internal/delivery/http/book` |
| IDs are created by storage when appropriate | Create input has no `ID`; repository returns `domain.Book` with generated `ID` | Client-provided create DTO is passed into SQL as a complete `domain.Book` with `ID` |

## Data Flow

### Create Book

| Step | Data Shape | Layer |
|---|---|---|
| Client sends JSON | HTTP create request DTO | `delivery/http/book` |
| Handler parses request | HTTP DTO | `delivery/http/book` |
| Handler calls application operation | Usecase create input without database-generated fields | `usecase/book` |
| Usecase applies business rules | Usecase input and domain concepts | `usecase/book`, `domain` |
| Repository inserts row | SQL parameters without client-provided `ID` | `infrastructure/postgres/book` |
| Database returns created row | `domain.Book` with generated `ID` | `infrastructure/postgres/book`, `domain` |
| Handler formats response | HTTP response DTO | `delivery/http/book` |

### Update Book

| Step | Data Shape | Layer |
|---|---|---|
| Client sends JSON | HTTP update request DTO, possibly partial | `delivery/http/book` |
| Handler parses request | HTTP DTO | `delivery/http/book` |
| Handler calls application operation | Usecase update input | `usecase/book` |
| Usecase decides update semantics | Usecase input and existing domain state if needed | `usecase/book` |
| Repository updates row | SQL parameters and row identifier | `infrastructure/postgres/book` |
| Database returns updated row | `domain.Book` | `infrastructure/postgres/book`, `domain` |
| Handler formats response | HTTP response DTO | `delivery/http/book` |

## Naming Guide

| Concept | Recommended Location | Purpose |
|---|---|---|
| `Book` | `internal/domain` | A book that exists in the system |
| `CreateBookRequest` | `internal/delivery/http/book` | JSON request body for HTTP create |
| `UpdateBookRequest` | `internal/delivery/http/book` | JSON request body for HTTP update |
| `BookResponse` | `internal/delivery/http/book` | JSON response shape for HTTP clients |
| `CreateBookInput` | `internal/usecase/book` | Application input for creating a book |
| `UpdateBookInput` | `internal/usecase/book` | Application input for updating a book |
| `Repository` | `internal/usecase/book` | Interface describing persistence needs of the usecase |
| `Repository` implementation | `internal/infrastructure/postgres/book` | Postgres-backed implementation of the usecase repository interface |

## Key Principle

Similar fields in different layers are not automatically duplication.

They represent different contracts:

| Contract | Example Concern |
|---|---|
| HTTP DTO | JSON names, optional request fields, public API shape |
| Usecase input | Business operation parameters |
| Domain entity | State and meaning inside the application |
| SQL row mapping | Database columns and persistence details |

Keep those contracts separate when they can change for different reasons.
