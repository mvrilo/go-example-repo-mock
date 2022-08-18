## go-example-repo-mock

[![CI](https://github.com/mvrilo/go-example-repo-mock/actions/workflows/ci.yaml/badge.svg)](https://github.com/mvrilo/go-example-repo-mock/actions/workflows/ci.yaml)

Example of application (data storage layer only) for a `User` domain-scoped service using Go, repository pattern for data access, github actions for ci and tests using `testify/suite` and `go-sqlmock`.

### Architecture:

- `model`
data structures mapping to the database schema and other domain data, e.g. definition of user, custom errors, etc

- `repository`
definition and implementation of storage access behavior, e.g. get data, save data, etc

- `mock`
data mocking for each repository implementation, e.g. fake sql results

### Testing

`make`

### Project structure

```
.
├── Makefile
├── README.md
├── go.mod
├── go.sum
├── mock                   # mock definitions
│   └── mysql              # by repository implementation
│       └── user.go
├── model                 # domain/database model definitions
│   └── user.go
└── repository             # repository interface
    ├── mysql              # and implementation
    │   ├── user.go
    │   └── user_test.go   # unit tests
    └── repository.go

5 directories, 9 files
```
