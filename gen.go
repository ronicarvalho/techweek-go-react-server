package gen

//go:generate go run .\cmd\tools\dotenv\main.go
//go:generate sqlc generate -f .\internal\store\pgstore\sqlc.yaml
//go:generate go mod tidy
