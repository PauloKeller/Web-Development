`run tests: go test ./... -coverprofile cover.out`
`remove mock files from cover.out: cat cover.out.tmp | grep -v "_mock.go" > cover.out`
`generate cover profile: go test ./... -coverprofile cover.out`
`open profile as html: go tool cover -html=cover.out`

`generate code from proto files: protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. *.proto`