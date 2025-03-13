# Julia API

### Run local

- Run `go run main.go`

### Build with docker-compose

- Run `docker-compose build`
- Run `docker-compose up`

### Running with docker
 - Run `docker exec -it julia_code python3 /app/julia.py --file TODO.MD`

### Commands

- Install API docs `go install github.com/swaggo/swag/cmd/swag@latest`
- Generate API docs `swag init --parseDependency --parseInternal`
- Run tests `go test`
- Add swag in path linux `export PATH=$(go env GOPATH)/bin:$PATH`
- Install goimports `go install golang.org/x/tools/cmd/goimports@latest`

### Copy to default project

- cp http/controllers/user_controller.go ../../julia/http/controllers/user_controller.go

### Task models


