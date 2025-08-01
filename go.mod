// go mod init <module_name>
// go mod tidy (checks if there is anything uninstalled in project/unused dependency)
// go install (global)/go get <module_name>

module github.com/Marie20767/go-playground

go 1.24.5

require github.com/stretchr/testify v1.10.0

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
