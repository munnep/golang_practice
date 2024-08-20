# some notes about the modules

# downloading and installing a module

- example gorilla mux for routing
  https://github.com/gorilla/mux

- Download and install
go get -u github.com/gorilla/mux

- check your go environment

go env

- check to the path GOPATH='/Users/patrick/go'

- here you will find all the downloads
/Users/patrick/go/pkg/mod/cache/download/github.com/gorilla/mux/@v

- use the following to make sure the correct modules are used and checked
go mod tidy

- verify all the modules
go mod version

- list all the modules
go list

- all packages used in your module
go list -m all

- check all the dependencies

go mod graph