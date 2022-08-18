# How to install
- copy .env.example to .env
- copy .air.toml.example to .air.toml.example
- make sure all the configuration is correct
- run go install github.com/cosmtrek/air@latest
- run go mod tidy

### Local Development
We use [air](https://github.com/cosmtrek/air) to enable hot reload on development
```
air
```

## Need to develop more
- Separate cmd, like migrate and seed
- Repository management
- Validation management