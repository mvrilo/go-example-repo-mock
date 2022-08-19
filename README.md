## go-example-repo-mock

[![CI](https://github.com/mvrilo/go-example-repo-mock/actions/workflows/ci.yaml/badge.svg)](https://github.com/mvrilo/go-example-repo-mock/actions/workflows/ci.yaml)

Example of application (data storage layer only) for a `User` domain-scoped service using Go, model definitions, github actions for ci and tests using `testify/suite` and `go-sqlmock`.

For a more decoupled architecture, check [the main branch](https://github.com/mvrilo/go-example-repo-mock/tree/main)

### Architecture:

- `model`

  - data structures mapping to the database schema and other domain data, e.g. definition of user, custom errors, etc
  - also contains the behavior for the storage access, e.g. get data, save data, etc
  - mock is used as a package inside the test module, e.g. user_mock_test.go with mock expectations from the tests

### Testing

`make`

### Project structure

```
.
├── Makefile
├── README.md
├── go.mod
├── go.sum
└── model
    ├── user.go
    ├── user_mock_test.go
    └── user_test.go

1 directory, 7 files
```
