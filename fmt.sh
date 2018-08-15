gofmt -w ./*.go

golint ./*.go

go vet ./*.go