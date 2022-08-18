# How to install
- copy .env.example to .env
- copy .air.toml.example to .air.toml.example
- make sure all the configuration is correct
- run go install github.com/cosmtrek/air@latest
- run go mod tidy
- run go install

### Local Development
We use [air](https://github.com/cosmtrek/air) to enable hot reload on development
```
air init
air
```
How to migrate
- using go run (preferred method for local development)
```
go run main.go migrate
[this method doesn't need to run 'go install' every modification in file/migration]
```
- using app_name (preferred method for server who need to run CI/CD everytime)
```
go-base-rest-api migrate
[this method need to run 'go install' every modification in migration]
```

## Need to develop more
- Seeder
