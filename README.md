## How to install
- copy .env.example to .env
- copy .env.test.example to .env.test
- make sure all env configuration is correct
- run go install github.com/cosmtrek/air@latest
- run go mod tidy
- run go install

## Local Development
We use [air](https://github.com/cosmtrek/air) to enable hot reload on development
```
air init
air
```
How to migrate and rollback
- using go run (preferred method for local development)
```
# Create Migration
go run main.go migrate --create=migration_name

# Migrate
go run main.go migrate

# Rollback
go run main.go migrate rollback
[this method doesn't need to run 'go install' every modification in file/migration]
```
- using app_name (preferred method for server who need to run CI/CD everytime)
```
# Create Migration
go-base-rest-api migrate --create=migration_name

# Migrate
go-base-rest-api migrate

# Rollback
go-base-rest-api migrate rollback
[this method need to run 'go install' every modification in migration]
```

Seeding
Keep in mind that seeding will delete all affected table data before inserting the seed
- using go run (preferred method for local development)
```
go run main.go seed
[this method doesn't need to run 'go install' every modification in file/migration]
```
- using app_name (preferred method for server who need to run CI/CD everytime)
```
go-base-rest-api seed
[this method need to run 'go install' every modification in migration]
```

## Testing
### Create test cases
All test case placed in tests folder, example in pkg/users/tests, pkg/auth/tests, etc
```go
// add this line below at the top of file
//go:build unit || integration || users || pkg || all
// Definition of that line above:
// Text after 'go:build' called as 'tags'
// unit is for unit test [make sure to chose only one unit or integration]
// integration is for integration test [make sure to chose only one unit or integration]
// users is for users pkg, it will make easier to run specific package test
// pkg is for pkg test, it will run all test case in pkg directory
// all tags is for all test case, it will run all test case in project

// IMPORTANT NOTE
// Feel free to use any tags you want, e.g. utils, main, migration, etc.

package tests
// code for test case below
```

### Run test cases
- Clean test cache before run test
```
go clean --testcache 
```
- Unit testing
```
 go test ./.../tests -tags=unit -v
```
- Integration testing
```
 go test ./.../tests -tags=integration -v
```
- Specific pkg test, example users pkg
```
 go test ./.../tests -tags=users -v
```
- All test case under pkg directory
```
 go test ./.../tests -tags=pkg -v
```
- All test case
```
 go test ./.../tests -tags=all -v
```

## Need to develop more
- DB Seeder
